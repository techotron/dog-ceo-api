import aiohttp
import asyncio
import json
import aiofiles
import os

from urllib.parse import urlparse

concurrency = 50
download_images = False
s3_bucket = "s3:/my_dog_bucket"


async def fetch_breeds(client):
    async with client.get("https://dog.ceo/api/breeds/list/all") as resp:
        return await resp.text()

async def fetch_breed_image_names(client, breed):
    async with client.get(f"https://dog.ceo/api/breed/{breed}/images") as resp:
        return await resp.text()

async def executor(client, queue):
    while True:
        url = await queue.get()
        async with client.get(url) as resp:
            if resp.status == 200:
                parsed_url = urlparse(url)
                filename = f"./images{parsed_url.path}"
                print(filename)
                os.makedirs(os.path.dirname(filename), exist_ok=True)
                image_file = await aiofiles.open(filename, mode="wb+")
                await image_file.write(await resp.read())
                await image_file.close()
        queue.task_done()

async def main():
    data_file = {}
    image_queue = asyncio.Queue()
    async with aiohttp.ClientSession() as client:
        breeds = await fetch_breeds(client)
        breeds_json = json.loads(breeds)
        dogs = breeds_json["message"]

        if download_images:
            workers = [
                asyncio.create_task(executor(client, image_queue)) for _ in range(concurrency)
            ]

        for breed in dogs:
            breed_images = await fetch_breed_image_names(client, breed)
            breed_image_json = json.loads(breed_images)
            for image in breed_image_json["message"]:
                _sub_breed = None
                url = urlparse(image)
                # print(f"{s3_bucket + url.path}")

                _, _, _breed, _file = url.path.split("/")

                if "-" in _breed:
                    _sub_breed = _breed.split("-")[1]
                    _breed = _breed.split("-")[0]

                
                if _breed in data_file:
                    if _sub_breed is not None:
                        if _sub_breed in data_file[_breed]:
                            data_file[_breed][_sub_breed] = data_file[_breed][_sub_breed] + [s3_bucket + url.path]
                        else:
                            data_file[_breed][_sub_breed] = [s3_bucket + url.path]

                else:
                    if _sub_breed is not None:
                        data_file[_breed] = {}
                        data_file[_breed][_sub_breed] = [s3_bucket + url.path]
                    else:
                        data_file[_breed] = [s3_bucket + url.path]

                if download_images:
                    print(f"adding {image} to queue")
                    await image_queue.put(image)

        if download_images:
            await image_queue.join()

            for worker in workers:
                worker.cancel()

    print(json.dumps(data_file))
         


loop = asyncio.get_event_loop()
loop.run_until_complete(main())
