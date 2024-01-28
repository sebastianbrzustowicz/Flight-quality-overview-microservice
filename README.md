# Filght data to PDF raport microservice

The program helps with converting given CSV data in a specific convention from a vehicle to PDF overview raport.    
Data refer to flight quality indicators of the quadcopter.    
This data is related to actual orientation of vehicle and its desired values defined in RPY angles.    
The integral quality indicators are computed and displayed.

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
The file must be sent to this endpoint for data handling:
| HTTP method | endpoint | description | request type | response type |
| -------------- | -------------- | -------------- | -------------- | -------------- |
| :yellow_circle: POST | /upload-data | get raport about flight data | csv file | pdf file |

# Example outcome
Data have to be .csv format and columns stands for:    
| rolld | pitchd | yawd | altituded | roll | pitch | yaw | altitude | isClamp |      
where "d" suffix stands for "desired".

### Input CSV

Example input `data.csv` as `file` keyword:
```
0.672918;0.869353;0.904524;0.612349;0.783116;0.545448;0.689145;0.234122;false
0.760927;0.352051;0.987521;0.310338;0.837838;0.740037;0.193987;0.018544;false
0.789194;0.998396;0.058161;0.241870;0.558699;0.378330;0.878447;0.083990;false
0.957174;0.124002;0.574128;0.278740;0.695367;0.184122;0.480707;0.107121;false
0.529210;0.374270;0.005429;0.288960;0.246190;0.476781;0.899539;0.093492;false
0.946348;0.151071;0.652267;0.801128;0.632496;0.242218;0.175317;0.043388;false
0.671881;0.964366;0.896576;0.400369;0.896163;0.349321;0.366211;0.685100;false
0.961425;0.883256;0.618784;0.029321;0.881504;0.672664;0.304091;0.256571;false
0.076191;0.012236;0.981465;0.968073;0.307531;0.228263;0.153329;0.927145;false
0.453702;0.255042;0.755973;0.782294;0.097271;0.752948;0.463828;0.225343;false
0.567253;0.804676;0.456280;0.576203;0.964866;0.946870;0.855016;0.120040;false
0.315552;0.716918;0.775555;0.939122;0.368368;0.272159;0.729856;0.063325;false
0.009080;0.134871;0.059352;0.752503;0.042761;0.903265;0.846345;0.498166;false
0.469579;0.648042;0.122009;0.688523;0.010712;0.195242;0.743968;0.416028;false
0.630682;0.080278;0.478924;0.524634;0.750729;0.681196;0.376879;0.534262;false
0.680669;0.664652;0.579580;0.188785;0.513167;0.330516;0.818094;0.329881;false
0.079654;0.356318;0.261963;0.683294;0.861316;0.802034;0.276226;0.847588;false
0.576071;0.534269;0.193146;0.576195;0.210927;0.831324;0.664772;0.652195;false
0.124588;0.347078;0.114943;0.133376;0.721130;0.233409;0.710036;0.102913;false
0.738029;0.173800;0.782023;0.305264;0.837455;0.550251;0.159438;0.462993;false
0.141238;0.862832;0.259256;0.677373;0.650469;0.898894;0.392166;0.122689;false
0.496114;0.524236;0.004074;0.149436;0.662996;0.162071;0.857503;0.571318;false
0.786533;0.050123;0.722229;0.956513;0.417124;0.234444;0.169866;0.756070;false
0.777083;0.751972;0.481983;0.231141;0.202641;0.628378;0.991647;0.633785;false
0.640587;0.559468;0.824754;0.988717;0.673812;0.942491;0.284695;0.998826;false
0.800520;0.119960;0.963167;0.222441;0.409590;0.025530;0.744000;0.865841;false
0.068057;0.781186;0.648432;0.823774;0.043830;0.542936;0.223669;0.898744;false
0.200393;0.983160;0.261191;0.462759;0.721616;0.488687;0.976053;0.046225;false
0.951640;0.670327;0.190188;0.850191;0.198508;0.438340;0.290641;0.098166;false
0.738567;0.362187;0.471556;0.121790;0.799507;0.681878;0.988370;0.444064;false
0.954744;0.961234;0.864682;0.476960;0.860916;0.009092;0.994702;0.950073;false
0.530719;0.235321;0.343663;0.506249;0.126145;0.579170;0.670768;0.444011;false
0.803176;0.818233;0.471648;0.725947;0.614754;0.579469;0.439062;0.952253;false
0.153555;0.554295;0.314646;0.685233;0.116096;0.775498;0.950191;0.900692;false
0.708828;0.048297;0.679930;0.017248;0.010065;0.265790;0.431165;0.136593;false

```

### Output PDF

The result consists of numerical results and plots.

<p align="center">
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/def60357-028f-47dc-936d-228869b19b1b" width="260" height="368" />
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/d0413cd0-345a-4467-9b17-34572b6e99d0" width="260" height="368"  />
  <img src="https://github.com/sebastianbrzustowicz/Flight-quality-overview-microservice/assets/66909222/58b71d04-645b-4d87-b5b6-c72538fe604a" width="260" height="368"  />
</p>

# License

Data-raport-microservice is released under the MIT license.

# Author

Sebastian Brzustowicz &lt;Se.Brzustowicz@gmail.com&gt;
