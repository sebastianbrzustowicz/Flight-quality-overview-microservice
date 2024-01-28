# Filght data to PDF raport microservice

The program helps with converting given CSV data in a specific convention from a vehicle to PDF overview raport.    
Data refer to flight quality indicators of the quadcopter.    
This data is related to actual orientation of vehicle and its desired values defined in RPY angles.    
The quality indicators (RMS and SSE) are computed and displayed, aswell as visualised raw data and control errors.

# Dockerization

Follow these simple commands to deploy application.
Building process:   
```console
docker build -t dataraport:go .
```
Running in background:
```console
docker run --name dataRaportContainer -d -p 8083:8083 dataraport:go
```
Now the containerized application should run properly.

# Endpoint
The CSV file must be sent to this endpoint for data handling as `file` keyword:
| HTTP method | endpoint | description | request type | response type |
| -------------- | -------------- | -------------- | -------------- | -------------- |
| :yellow_circle: POST | /upload-data | get raport about flight data | CSV file | PDF file |

# Example outcome
Data have to be .csv format and columns stands for:    
| rolld | pitchd | yawd | altituded | roll | pitch | yaw | altitude | isClamp |      
where "d" suffix stands for "desired".

### Input CSV

Example data sample can be found in `data.csv` file in the project folder.    

Example input:
```
0.000000;0.200000;0.000000;20.000000;0.000000;0.200000;0.000000;20.000000;false
0.300000;0.000000;0.174524;25.000000;0.032357;0.185000;0.050000;20.500000;false
0.300000;0.000000;0.342020;25.000000;0.082606;0.170000;0.174524;21.000000;false
0.300000;0.000000;0.500000;25.000000;0.130000;0.140000;0.342020;22.000000;false
0.300000;0.000000;0.642788;25.000000;0.200000;0.100000;0.500000;23.000000;false
0.300000;0.000000;0.766044;25.000000;0.250000;0.060000;0.642788;24.000000;false
0.300000;0.000000;0.866025;25.000000;0.280000;0.030000;0.766044;24.500000;false
0.300000;0.000000;0.939693;25.000000;0.290000;0.020000;0.866025;24.800000;false
0.300000;0.000000;0.984808;25.000000;0.300000;0.010000;0.939693;25.000000;false
0.300000;0.000000;1.000000;25.000000;0.300000;0.000000;0.984808;25.000000;false
0.300000;0.000000;1.000000;25.000000;0.300000;0.000000;1.000000;25.000000;false
0.300000;0.000000;1.000000;25.000000;0.300000;0.000000;1.000000;25.000000;false
...

```

### Output PDF

The result consists of numerical results and plots.

<p align="center">
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/323493fa-1cfb-408b-bda5-b4e0d1d0e03a" width="260" height="368" />
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/765507c1-a3d3-4f9e-9482-ba86a8945205" width="260" height="368"  />
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/44dd54bd-703b-4d0f-b14b-8a8ab061780f" width="260" height="368"  />
</p>

# License

Flight-quality-overview-microservice is released under the MIT license.

# Author

Sebastian Brzustowicz &lt;Se.Brzustowicz@gmail.com&gt;
