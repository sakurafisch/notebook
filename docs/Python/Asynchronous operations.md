# Asynchronous operations

## 传统的写法导致阻塞

```pyhon
from timeit import default_timer
import requests

def load_data(delay):
    print(f'Starting {delay} second timer')
    text = requests.get(f'https://httpbin.org/delay/{delay}').text
    print(f'Completed {delay} second timer')
    return text

def run_demo():
    start_time = default_timer()

    two_data = load_data(2)
    three_data = load_data(3)

    elapsed_time = default_timer() - start_time
    print(f'The operation took {elapsed_time:.2} seconds')

def main():
    run_demo()

if __name__ == "__main__":
    main()
```

## 重构为异步代码

```python
# require Python 3.7 or higher
from timeit import default_timer
import aiohttp
import asyncio

async def load_data(session, delay):
    print(f'Starting {delay} second timer')
    async with session.get(f'https://httpbin.org/delay/{delay}') as resp:
        text = await resp.text()
        print(f'Completed {delay} second timer')
        return text

async def main():
    # Start the timer
    start_time = default_timer()

    # Create a single session
    async with aiohttp.ClientSession() as session:
        # Setup our task
        two_task = asyncio.create_task(load_data(session, 2))
        three_task = asyncio.create_task(load_data(session, 3))
    
        # Simulate other processing
        await asyncio.sleep(1)
        print('Doing other work')

        # Let's go get our values
        two_result = await two_task
        three_result = await three_task
        
        # Print our result
        elapsed_time = default_timer() - start_time
        print(f'The operation took {elapsed_time:.2} seconds')



if __name__ == "__main__":
    asyncio.run(main())

```

在上文用到的`aiohttp`可参考[这里](https://docs.aiohttp.org/en/stable/client_quickstart.html#response-content-and-status-code)

