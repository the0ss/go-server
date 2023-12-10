## Sorting Server

Simple API server for sorting 2D array.

### Running on your system

- Install Docker
- Run this command ( This will download the image from hub and run locally on your system on Port 8000)

```
docker run -p 8000:8000 the0ss/sorting-server
```

### Two endpoints

- Sequential sorting

```
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]}' http://localhost:8000/process-single
```

- Concurrent sorting

```
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]}' http://localhost:8000/process-concurrent
```

- If using Postman
  - Sequential(POST)
  ```
  http://localhost:8000/process-single
  ```
  - Concurrent(POST)
  ```
  http://localhost:8000/process-concurrent
  ```
  - Body Remain same in both
  ```
  {
  "to_sort": [[1, 3, 2], [4, 7, 6], [7, 8, 9], [7, 8, 10], [11, 8, 9], [7, 2, 9]]
  }
  ```

### Deployed On render

- cURL
```
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]}' https://sorting-server-rjv1.onrender.com/process-single
```
```
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 2, 1], [6, 5, 4], [9, 8, 7]]}' https://sorting-server-rjv1.onrender.com/process-concurrent
```