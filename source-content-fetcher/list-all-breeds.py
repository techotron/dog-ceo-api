import aiohttp
import asyncio

import json

concurrency = 50

async def fetch_breeds(client):
    async with client.get("https://dog.ceo/api/breeds/list/all") as resp:
        return await resp.text()


async def fetch_sub_breed_image_url(client, breed, sub_breed):
    async with client.get(f"https://dog.ceo/api/breed/{breed}/{sub_breed}/images") as resp:
        return await resp.text()


async def fetch_breed_image_url(client, breed):
    async with client.get(f"https://dog.ceo/api/breed/{breed}/images") as resp:
        return await resp.text()


async def executor(client, queue):
    while True:
        breed, sub_breed = await queue.get()

        if sub_breed:
            sub_breed_images = await fetch_sub_breed_image_url(client, breed, sub_breed)
            sub_breed_images_json = json.loads(sub_breed_images)
            for url in sub_breed_images_json["message"]:
                print(f"('{breed}', '{sub_breed}', '{url}'),")
        else:
            breed_images = await fetch_breed_image_url(client, breed)
            breed_images_json = json.loads(breed_images)
            for url in breed_images_json["message"]:
                print(f"('{breed}', NULL, '{url}'),")

        queue.task_done()

async def main():
    get_images_for_breeds = []
    all_images = []
    queue = asyncio.Queue()
    async with aiohttp.ClientSession() as client:
        breeds = await fetch_breeds(client)
        breeds_json = json.loads(breeds)
        dogs = breeds_json["message"]

        workers = [
            asyncio.create_task(executor(client, queue)) for _ in range(concurrency)
        ]

        for breed in dogs:
            # Insert Breeds
            # print(f"('{breed}'),")

            # Insert Sub Breeds
            # if dogs[breed]:
            #     for sub_breed in dogs[breed]:
            #         print(f"('{breed}', '{sub_breed}'),")

            # Insert Images
            if dogs[breed]:
                for sub_breed in dogs[breed]:
                    get_images_for_breeds.append((breed, sub_breed))
                    await queue.put((breed, sub_breed))
                    # print(f"((SELECT breed_id FROM breeds WHERE breed_name = '{breed}'), (SELECT sub_breed_id FROM sub_breeds WHERE sub_breed_name = '{sub_breed}'), 'http://pic.jpeg'),")
            else:
                get_images_for_breeds.append((breed, None))
                await queue.put((breed, None))


        await queue.join()

        for worker in workers:
            worker.cancel()


loop = asyncio.get_event_loop()
loop.run_until_complete(main())
