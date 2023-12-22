### API接口文档

##### 创建推理服务

- URL:   /api/service

- Method: post

-  参数：

| 字段     |           |        | 类型             | 说明                                                    |
| -------- | --------- | ------ | ---------------- | ------------------------------------------------------- |
| image    |           |        | 字符串           | 镜像地址                                                |
| cpu      |           |        | Int32            | CPU大小 核                                              |
| gpu      |           |        | Int32            | gpu卡数                                                 |
| memory   |           |        | Int32            | 内存 单位G                                              |
| port     |           |        | Int32            | 端口                                                    |
| storages |           |        | 存储挂载数组     |                                                         |
|          | mountPath |        | 项目指定路径地址 |                                                         |
|          | nfs       |        |                  | 其中nfs的地址（其中nfs和miniio其中选一个节点就可以）    |
|          |           | path   |                  |                                                         |
|          |           | server |                  |                                                         |
|          | miniio    |        |                  | （需要配合默认的miniIo地址）                            |
|          |           | path   | 桶命/路径        | 其中miniio的地址（其中nfs和miniio其中选一个节点就可以） |
|          |           |        |                  |                                                         |


eg：


  ```json
  {
    "source":"aog",
    "image": "harbor.dm-ai.com/model/serving:v0",  
    "cpu": 1, 
    "memory": 1,
    "port": 8089,
    "storages": [
      {
        "mountPath": "/data/project/models/",
        "nfs": {
          "path": "/data/nfsshare/serving-data/test/models/",
          "server": "10.12.16.41"
        }
      }
      },
    ]
  }
  ```

- 结果

  ```json
  {
    "code": 0,
    "error": "",
     "data":{
      "url": "http://10.12.16.40:32509",
      "name": "service-20230724170708-xvlb"
     }
  }  
  ```

##### 删除推理服务

- URL:/api/service/:servicename

- Method: delete

- 结果

  ```json
  {
  "code": 0,
  "error": "",
  "data": true
  }
  ```







##### 查看推理服务状态

- URL:/api/service/status/::serviceName

- Method: get

- 结果

  ```json
  {
  "code": 0,
  "error": "",
      "data":{
        "name": "service-20230725114336-rxob",
        "status": "Running"
      }
  }
  ```



##### 查看推理服务状态列表

- URL:/api/service/status/list

- Method: post

- 参数

- ```json
  [
    "service-20230725175047-slmh"
  ]
  ```

  

- 结果

  ```json
  {
  "code": 0,
  "error": "",
  "data":[
      {
        "name": "service-20230725175047-slmh",
        "status": "Pending"
      }
  ]
  }
  ```





##### 查看推理服务日志

- URL:/api/service/log/:serviceName

- Method: get

- 参数

- | 字段 | 类型 | 说明                      |
  | ---- | ---- | ------------------------- |
  | line | int  | 最多返回日志行数 默认1000 |


- 结果

  ```json
  
  
  {
    "code": 0,
    "error": "",
    "data": "51/469 [00:09<01:23,  5.01it/s]\rTrain Epoch #0:  11%|█         | 52/469 [00:09<01:21,  5.11it/s]\rTrain Epoch #0:  11%|█▏        | 53/469 [00:10<01:22,  5.07it/s]\rTrain Epoch #0:  12%|█▏        | 54/469 [00:10<01:11,  5.79it/s]\rTrain Epoch #0:  12%|█▏        | 55/469 [00:10<01:15,  5.46it/s]\rTrain Epoch #0:  12%|█▏        | 56/469 [00:10<01:17,  5.35it/s]\rTrain Epoch #0:  12%|█▏        | 57/469 [00:10<01:19,  5.20it/s]\rTrain Epoch #0:  12%|█▏        | 58/469 [00:11<01:17,  5.33it/s]\rTrain Epoch #0:  13%|█▎        | 59/469 [00:11<01:18,  5.19it/s]\rTrain Epoch #0:  13%|█▎        | 60/469 [00:11<01:18,  5.18it/s]\rTrain Epoch #0:  13%|█▎        | 61/469 [00:11<01:08,  5.94it/s]\rTrain Epoch #0:  13%|█▎        | 62/469 [00:11<01:13,  5.55it/s]\rTrain Epoch #0:  13%|█▎        | 63/469 [00:11<01:15,  5.40it/s]\rTrain Epoch #0:  14%|█▎        | 64/469 [00:12<01:15,  5.36it/s]\rTrain Epoch #0:  14%|█▍        | 65/469 [00:12<01:16,  5.27it/s]\rTrain Epoch #0:  14%|█▍        | 66/469 [00:12<01:17,  5.17it/s]\rTrain Epoch #0:  14%|█▍        | 67/469 [00:12<01:19,  5.06it/s]\rTrain Epoch #0:  14%|█▍        | 68/469 [00:12<01:08,  5.87it/s]\rTrain Epoch #0:  15%|█▍        | 69/469 [00:13<01:12,  5.53it/s]\rTrain Epoch #0:  15%|█▍        | 70/469 [00:13<01:12,  5.49it/s]\rTrain Epoch #0:  15%|█▌        | 71/469 [00:13<01:15,  5.29it/s]\rTrain Epoch #0:  15%|█▌        | 72/469 [00:13<01:15,  5.23it/s]\rTrain Epoch #0:  16%|█▌        | 73/469 [00:13<01:17,  5.09it/s]\rTrain Epoch #0:  16%|█▌        | 74/469 [00:14<01:16,  5.14it/s]\rTrain Epoch #0:  16%|█▌        | 75/469 [00:14<01:16,  5.17it/s]\rTrain Epoch #0:  16%|█▌        | 76/469 [00:14<01:06,  5.93it/s]\rTrain Epoch #0:  16%|█▋        | 77/469 [00:14<01:10,  5.53it/s]\rTrain Epoch #0:  17%|█▋        | 78/469 [00:14<01:12,  5.39it/s]\rTrain Epoch #0:  17%|█▋        | 79/469 [00:14<01:14,  5.26it/s]\rTrain Epoch #0:  17%|█▋        | 80/469 [00:15<01:13,  5.29it/s]\rTrain Epoch #0:  17%|█▋        | 81/469 [00:15<01:16,  5.10it/s]\rTrain Epoch #0:  17%|█▋        | 82/469 [00:15<01:15,  5.14it/s]\rTrain Epoch #0:  18%|█▊        | 83/469 [00:15<01:13,  5.23it/s]\rTrain Epoch #0:  18%|█▊        | 84/469 [00:15<01:06,  5.82it/s]\rTrain Epoch #0:  18%|█▊        | 85/469 [00:16<01:09,  5.52it/s]\rTrain Epoch #0:  18%|█▊        | 86/469 [00:16<01:10,  5.40it/s]\rTrain Epoch #0:  19%|█▊        | 87/469 [00:16<01:11,  5.37it/s]\rTrain Epoch #0:  19%|█▉        | 88/469 [00:16<01:12,  5.24it/s]\rTrain Epoch #0:  19%|█▉        | 89/469 [00:16<01:14,  5.12it/s]\rTrain Epoch #0:  19%|█▉        | 90/469 [00:16<01:03,  6.00it/s]\rTrain Epoch #0:  19%|█▉        | 91/469 [00:17<01:14,  5.09it/s]\rTrain Epoch #0:  20%|█▉        | 92/469 [00:17<01:08,  5.54it/s]\rTrain Epoch #0:  20%|█▉        | 93/469 [00:17<01:09,  5.40it/s]\rTrain Epoch #0:  20%|██        | 94/469 [00:17<01:11,  5.28it/s]\rTrain Epoch #0:  20%|██        | 95/469 [00:17<01:11,  5.23it/s]\rTrain Epoch #0:  20%|██        | 96/469 [00:18<01:13,  5.08it/s]\rTrain Epoch #0:  21%|██        | 97/469 [00:18<01:12,  5.16it/s]\rTrain Epoch #0:  21%|██        | 98/469 [00:18<01:13,  5.02it/s]\rTrain Epoch #0:  21%|██        | 99/469 [00:18<01:12,  5.08it/s]\rTrain Epoch #0:  21%|██▏       | 100/469 [00:18<01:02,  5.90it/s]INFO: {'lr': 0.0070362473347547975, 'train': {'Accuracy': 28.712871287128714, 'loss': 4.41335366661948}}\n\rTrain Epoch #0:  22%|██▏       | 101/469 [00:19<01:06,  5.56it/s]\rTrain Epoch #0:  22%|██▏       | 102/469 [00:19<01:07,  5.45it/s]\rTrain Epoch #0:  22%|██▏       | 103/469 [00:19<01:08,  5.37it/s]\rTrain Epoch #0:  22%|██▏       | 104/469 [00:19<01:08,  5.30it/s]\rTrain Epoch #0:  22%|██▏       | 105/469 [00:19<01:08,  5.28it/s]\rTrain Epoch #0:  23%|██▎       | 106/469 [00:19<01:01,  5.89it/s]\rTrain Epoch #0:  23%|██▎       | 107/469 [00:20<01:04,  5.65it/s]\rTrain Epoch #0:  23%|██▎       | 108/469 [00:20<01:06,  5.46it/s]\rTrain Epoch #0:  23%|██▎       | 109/469 [00:20<01:08,  5.26it/s]\rTrain Epoch #0:  23%|██▎       | 110/469 [00:20<01:06,  5.36it/s]\rTrain Epoch #0:  24%|██▎       | 111/469 [00:20<01:06,  5.36it/s]\rTrain Epoch #0:  24%|██▍       | 112/469 [00:21<01:12,  4.95it/s]\rTrain Epoch #0:  24%|██▍       | 114/469 [00:21<01:06,  5.31it/s]\rTrain Epoch #0:  25%|██▍       | 115/469 [00:21<01:07,  5.25it/s]\rTrain Epoch #0:  25%|██▍       | 116/469 [00:21<01:08,  5.18it/s]\rTrain Epoch #0:  25%|██▍       | 117/469 [00:22<01:07,  5.22it/s]\rTrain Epoch #0:  25%|██▌       | 118/469 [00:22<01:07,  5.19it/s]\rTrain Epoch #0:  25%|██▌       | 119/469 [00:22<01:08,  5.11it/s]\rTrain Epoch #0:  26%|██▌       | 120/469 [00:22<01:09,  5.03it/s]\rTrain Epoch #0:  26%|██▌       | 122/469 [00:22<01:03,  5.43it/s]\rTrain Epoch #0:  26%|██▌       | 123/469 [00:23<01:05,  5.27it/s]\rTrain Epoch #0:  26%|██▋       | 124/469 [00:23<01:07,  5.11it/s]\rTrain Epoch #0:  27%|██▋       | 125/469 [00:23<01:07,  5.13it/s]\rTrain Epoch #0:  27%|██▋       | 126/469 [00:23<01:06,  5.18it/s]\rTrain Epoch #0:  27%|██▋       | 127/469 [00:23<01:08,  4.96it/s]\rTrain Epoch #0:  27%|██▋       | 128/469 [00:24<01:07,  5.07it/s]\rTrain Epoch #0:  28%|██▊       | 129/469 [00:24<01:08,  4.96it/s]\rTrain Epoch #0:  28%|██▊       | 130/469 [00:24<01:06,  5.13it/s]\rTrain Epoch #0:  28%|██▊       | 131/469 [00:24<01:08,  4.93it/s]\rTrain Epoch #0:  28%|██▊       | 132/469 [00:24<01:06,  5.05it/s]\rTrain Epoch #0:  28%|██▊       | 133/469 [00:25<01:07,  5.01it/s]\rTrain Epoch #0:  29%|██▊       | 134/469 [00:25<01:06,  5.07it/s]\rTrain Epoch #0:  29%|██▉       | 135/469 [00:25<01:07,  4.98it/s]\rTrain Epoch #0:  29%|██▉       | 136/469 [00:25<01:05,  5.09it/s]\rTrain Epoch #0:  29%|██▉       | 137/469 [00:25<01:06,  4.96it/s]\rTrain Epoch #0:  29%|██▉       | 138/469 [00:26<01:06,  4.95it/s]\rTrain Epoch #0:  30%|██▉       | 139/469 [00:26<01:06,  4.97it/s]\rTrain Epoch #0:  30%|██▉       | 140/469 [00:26<01:04,  5.11it/s]\rTrain Epoch #0:  30%|███       | 141/469 [00:26<01:05,  4.99it/s]\rTrain Epoch #0:  30%|███       | 142/469 [00:26<01:04,  5.05it/s]\rTrain Epoch #0:  30%|███       | 143/469 [00:27<01:05,  4.94it/s]\rTrain Epoch #0:  31%|███       | 144/469 [00:27<01:03,  5.10it/s]\rTrain Epoch #0:  31%|███       | 145/469 [00:27<01:04,  5.04it/s]\rTrain Epoch #0:  31%|███       | 146/469 [00:27<01:05,  4.97it/s]\rTrain Epoch #0:  31%|███▏      
  }
  ```



##### 创建训练任务

- URL:/api/train

- Method: post

- 参数：

| 字段     |           |        | 类型             | 说明                                                    |
| -------- | --------- | ------ | ---------------- | ------------------------------------------------------- |
| image    |           |        | 字符串           | 镜像地址                                                |
| cpu      |           |        | Int32            | CPU大小 核                                              |
| gpu      |           |        | Int32            | gpu卡数                                                 |
| memory   |           |        | Int32            | 内存单位 G                                              |
| command  |           |        | 字符串数组       |                                                         |
| storages |           |        | 存储挂载数组     |                                                         |
|          | mountPath |        | 项目指定路径地址 |                                                         |
|          | nfs       |        |                  | 其中nfs的地址（其中nfs和miniio其中选一个节点就可以）    |
|          |           | path   |                  |                                                         |
|          |           | server |                  |                                                         |
|          | miniio    |        |                  | （需要配合默认的miniIo地址）                            |
|          |           | path   | 桶命/路径        | 其中miniio的地址（其中nfs和miniio其中选一个节点就可以） |
|          |           |        |                  |                                                         |


eg：


  ```json
{
  "image": "harbor.dm-ai.com/model/train:v0",
  "cpu": "8",
  "memory": "16G",
  "command": [
    "python",
    "example.py"
  ],
  "storages": [
    {
      "mountPath": "/app1/data/",
      "nfs": {
        "path": "/data/nfsshare/train-test/data/",
        "server": "10.12.16.41"
      }
    },
    {
      "mountPath": "/app1/work_space/",
      "nfs": {
        "path": "//data/nfsshare/train-test/output/",
        "server": "10.12.16.41"
      }
    }
  ]
}
  ```

- 结果

  ```json
  {
  "code": 0,
  "error": "",
      "data":{
      "name": "job-20230725112049-yxce"
     }
  }
  ```

##### 删除训练任务

- URL:/api/train/:trainName

- Method: delete

- 结果

  ```json
  {
  "code": 0,
  "error": "",
  "data": true
  }
  ```

##### 查看任务训练状态

- URL:/api/train/status/::trainName

- Method: get

- 结果

  ```json
  {
  "code": 0,
  "error": "",
      "data":{
        "name": "job-20230725114336-rxob",
        "status": "Running"
      }
  }
  ```

##### 查看任务训练状态列表

- URL:/api/train/status/list

- Method: post

- 结果

  ```json
  {
  "code": 0,
  "error": "",
  "data":[
      {
        "name": "job-20230725175047-slmh",
        "status": "Pending"
      }
  ]
  }
  ```

##### 查看训练任务日志

- URL:/api/train/log/:trainName

- Method: get

- 参数

- | 字段 | 类型 | 说明                      |
  | ---- | ---- | ------------------------- |
  | line | int  | 最多返回日志行数 默认1000 |


- 结果

  ```json
  
  {
    "code": 0,
    "error": "",
    "data": " 51/469 [00:09<01:23,  5.01it/s]\rTrain Epoch #0:  11%|█         | 52/469 [00:09<01:21,  5.11it/s]\rTrain Epoch #0:  11%|█▏        | 53/469 [00:10<01:22,  5.07it/s]\rTrain Epoch #0:  12%|█▏        | 54/469 [00:10<01:11,  5.79it/s]\rTrain Epoch #0:  12%|█▏        | 55/469 [00:10<01:15,  5.46it/s]\rTrain Epoch #0:  12%|█▏        | 56/469 [00:10<01:17,  5.35it/s]\rTrain Epoch #0:  12%|█▏        | 57/469 [00:10<01:19,  5.20it/s]\rTrain Epoch #0:  12%|█▏        | 58/469 [00:11<01:17,  5.33it/s]\rTrain Epoch #0:  13%|█▎        | 59/469 [00:11<01:18,  5.19it/s]\rTrain Epoch #0:  13%|█▎        | 60/469 [00:11<01:18,  5.18it/s]\rTrain Epoch #0:  13%|█▎        | 61/469 [00:11<01:08,  5.94it/s]\rTrain Epoch #0:  13%|█▎        | 62/469 [00:11<01:13,  5.55it/s]\rTrain Epoch #0:  13%|█▎        | 63/469 [00:11<01:15,  5.40it/s]\rTrain Epoch #0:  14%|█▎        | 64/469 [00:12<01:15,  5.36it/s]\rTrain Epoch #0:  14%|█▍        | 65/469 [00:12<01:16,  5.27it/s]\rTrain Epoch #0:  14%|█▍        | 66/469 [00:12<01:17,  5.17it/s]\rTrain Epoch #0:  14%|█▍        | 67/469 [00:12<01:19,  5.06it/s]\rTrain Epoch #0:  14%|█▍        | 68/469 [00:12<01:08,  5.87it/s]\rTrain Epoch #0:  15%|█▍        | 69/469 [00:13<01:12,  5.53it/s]\rTrain Epoch #0:  15%|█▍        | 70/469 [00:13<01:12,  5.49it/s]\rTrain Epoch #0:  15%|█▌        | 71/469 [00:13<01:15,  5.29it/s]\rTrain Epoch #0:  15%|█▌        | 72/469 [00:13<01:15,  5.23it/s]\rTrain Epoch #0:  16%|█▌        | 73/469 [00:13<01:17,  5.09it/s]\rTrain Epoch #0:  16%|█▌        | 74/469 [00:14<01:16,  5.14it/s]\rTrain Epoch #0:  16%|█▌        | 75/469 [00:14<01:16,  5.17it/s]\rTrain Epoch #0:  16%|█▌        | 76/469 [00:14<01:06,  5.93it/s]\rTrain Epoch #0:  16%|█▋        | 77/469 [00:14<01:10,  5.53it/s]\rTrain Epoch #0:  17%|█▋        | 78/469 [00:14<01:12,  5.39it/s]\rTrain Epoch #0:  17%|█▋        | 79/469 [00:14<01:14,  5.26it/s]\rTrain Epoch #0:  17%|█▋        | 80/469 [00:15<01:13,  5.29it/s]\rTrain Epoch #0:  17%|█▋        | 81/469 [00:15<01:16,  5.10it/s]\rTrain Epoch #0:  17%|█▋        | 82/469 [00:15<01:15,  5.14it/s]\rTrain Epoch #0:  18%|█▊        | 83/469 [00:15<01:13,  5.23it/s]\rTrain Epoch #0:  18%|█▊        | 84/469 [00:15<01:06,  5.82it/s]\rTrain Epoch #0:  18%|█▊        | 85/469 [00:16<01:09,  5.52it/s]\rTrain Epoch #0:  18%|█▊        | 86/469 [00:16<01:10,  5.40it/s]\rTrain Epoch #0:  19%|█▊        | 87/469 [00:16<01:11,  5.37it/s]\rTrain Epoch #0:  19%|█▉        | 88/469 [00:16<01:12,  5.24it/s]\rTrain Epoch #0:  19%|█▉        | 89/469 [00:16<01:14,  5.12it/s]\rTrain Epoch #0:  19%|█▉        | 90/469 [00:16<01:03,  6.00it/s]\rTrain Epoch #0:  19%|█▉        | 91/469 [00:17<01:14,  5.09it/s]\rTrain Epoch #0:  20%|█▉        | 92/469 [00:17<01:08,  5.54it/s]\rTrain Epoch #0:  20%|█▉        | 93/469 [00:17<01:09,  5.40it/s]\rTrain Epoch #0:  20%|██        | 94/469 [00:17<01:11,  5.28it/s]\rTrain Epoch #0:  20%|██        | 95/469 [00:17<01:11,  5.23it/s]\rTrain Epoch #0:  20%|██        | 96/469 [00:18<01:13,  5.08it/s]\rTrain Epoch #0:  21%|██        | 97/469 [00:18<01:12,  5.16it/s]\rTrain Epoch #0:  21%|██        | 98/469 [00:18<01:13,  5.02it/s]\rTrain Epoch #0:  21%|██        | 99/469 [00:18<01:12,  5.08it/s]\rTrain Epoch #0:  21%|██▏       | 100/469 [00:18<01:02,  5.90it/s]INFO: {'lr': 0.0070362473347547975, 'train': {'Accuracy': 28.712871287128714, 'loss': 4.41335366661948}}\n\rTrain Epoch #0:  22%|██▏       | 101/469 [00:19<01:06,  5.56it/s]\rTrain Epoch #0:  22%|██▏       | 102/469 [00:19<01:07,  5.45it/s]\rTrain Epoch #0:  22%|██▏       | 103/469 [00:19<01:08,  5.37it/s]\rTrain Epoch #0:  22%|██▏       | 104/469 [00:19<01:08,  5.30it/s]\rTrain Epoch #0:  22%|██▏       | 105/469 [00:19<01:08,  5.28it/s]\rTrain Epoch #0:  23%|██▎       | 106/469 [00:19<01:01,  5.89it/s]\rTrain Epoch #0:  23%|██▎       | 107/469 [00:20<01:04,  5.65it/s]\rTrain Epoch #0:  23%|██▎       | 108/469 [00:20<01:06,  5.46it/s]\rTrain Epoch #0:  23%|██▎       | 109/469 [00:20<01:08,  5.26it/s]\rTrain Epoch #0:  23%|██▎       | 110/469 [00:20<01:06,  5.36it/s]\rTrain Epoch #0:  24%|██▎       | 111/469 [00:20<01:06,  5.36it/s]\rTrain Epoch #0:  24%|██▍       | 112/469 [00:21<01:12,  4.95it/s]\rTrain Epoch #0:  24%|██▍       | 114/469 [00:21<01:06,  5.31it/s]\rTrain Epoch #0:  25%|██▍       | 115/469 [00:21<01:07,  5.25it/s]\rTrain Epoch #0:  25%|██▍       | 116/469 [00:21<01:08,  5.18it/s]\rTrain Epoch #0:  25%|██▍       | 117/469 [00:22<01:07,  5.22it/s]\rTrain Epoch #0:  25%|██▌       | 118/469 [00:22<01:07,  5.19it/s]\rTrain Epoch #0:  25%|██▌       | 119/469 [00:22<01:08,  5.11it/s]\rTrain Epoch #0:  26%|██▌       | 120/469 [00:22<01:09,  5.03it/s]\rTrain Epoch #0:  26%|██▌       | 122/469 [00:22<01:03,  5.43it/s]\rTrain Epoch #0:  26%|██▌       | 123/469 [00:23<01:05,  5.27it/s]\rTrain Epoch #0:  26%|██▋       | 124/469 [00:23<01:07,  5.11it/s]\rTrain Epoch #0:  27%|██▋       | 125/469 [00:23<01:07,  5.13it/s]\rTrain Epoch #0:  27%|██▋       | 126/469 [00:23<01:06,  5.18it/s]\rTrain Epoch #0:  27%|██▋       | 127/469 [00:23<01:08,  4.96it/s]\rTrain Epoch #0:  27%|██▋       | 128/469 [00:24<01:07,  5.07it/s]\rTrain Epoch #0:  28%|██▊       | 129/469 [00:24<01:08,  4.96it/s]\rTrain Epoch #0:  28%|██▊       | 130/469 [00:24<01:06,  5.13it/s]\rTrain Epoch #0:  28%|██▊       | 131/469 [00:24<01:08,  4.93it/s]\rTrain Epoch #0:  28%|██▊       | 132/469 [00:24<01:06,  5.05it/s]\rTrain Epoch #0:  28%|██▊       | 133/469 [00:25<01:07,  5.01it/s]\rTrain Epoch #0:  29%|██▊       | 134/469 [00:25<01:06,  5.07it/s]\rTrain Epoch #0:  29%|██▉       | 135/469 [00:25<01:07,  4.98it/s]\rTrain Epoch #0:  29%|██▉       | 136/469 [00:25<01:05,  5.09it/s]\rTrain Epoch #0:  29%|██▉       | 137/469 [00:25<01:06,  4.96it/s]\rTrain Epoch #0:  29%|██▉       | 138/469 [00:26<01:06,  4.95it/s]\rTrain Epoch #0:  30%|██▉       | 139/469 [00:26<01:06,  4.97it/s]\rTrain Epoch #0:  30%|██▉       | 140/469 [00:26<01:04,  5.11it/s]\rTrain Epoch #0:  30%|███       | 141/469 [00:26<01:05,  4.99it/s]\rTrain Epoch #0:  30%|███       | 142/469 [00:26<01:04,  5.05it/s]\rTrain Epoch #0:  30%|███       | 143/469 [00:27<01:05,  4.94it/s]\rTrain Epoch #0:  31%|███       | 144/469 [00:27<01:03,  5.10it/s]\rTrain Epoch #0:  31%|███       | 145/469 [00:27<01:04,  5.04it/s]\rTrain Epoch #0:  31%|███       | 146/469 [00:27<01:05,  4.97it/s]\rTrain Epoch #0:  31%|███▏      
  }
  ```
