from fastapi import FastAPI
import pygeohash as Geohash
import random
random.seed(5)
app = FastAPI()
givenDistance = 100.0

def calculateDistance(pos1,pos2):
    return ((pos1[0]-pos2[0])**2 + (pos1[1]-pos2[1])**2)


drivers = []
for i in range(10):
    temp = {}
    x = random.random()*10
    y = random.random()*10
    temp['geoHash'] = Geohash.encode(x,y)
    temp['Name'] = 'Name' + str(i)
    temp['DId'] = i
    drivers.append(temp)
    

@app.post("/get_nearby_drivers")
async def getNearbyDrivers(x:float, y:float):
    position = (x,y)
    ret = []
    for d in drivers:
        if calculateDistance(position, Geohash.decode(d['geoHash']))<givenDistance:
            ret.append(d)
    return(ret)
