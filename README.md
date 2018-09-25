# seatsDistributionSystem

Test application listens to port `:8080`. Commands to run application from `cmd` directory:
```
sudo docker build -t seatsdistributionsystem .
sudo docker run --publish 8080:8080 --name test --rm seatsdistributionsystem
```

## API:


### GetNewSalon returns new aircraft salon with provided dimensions and salon ID
Request: 
```
localhost:8080/salon/new
{
	"Rows": 2,
	"Blocks": 2,
	"SeatCount":[2,2]
}

```
Response:
```
{
    "SalonID": "188828f8-1782-4758-9c8f-4bb00bda5ed3",
    "Salon": {
        "Rows": [
            {
                "Blocks": [
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "1A",
                                "SeatClass": "window"
                            },
                            {
                                "Assigned": false,
                                "Num": "1B",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "1C",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "1D",
                                "SeatClass": "window"
                            }
                        ]
                    }
                ]
            },
            {
                "Blocks": [
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "2A",
                                "SeatClass": "window"
                            },
                            {
                                "Assigned": false,
                                "Num": "2B",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "2C",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "2D",
                                "SeatClass": "window"
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
```
### GetSalon returns aircraft salon by ID
Request: 
```
localhost:8080/salon/get
{
	"SalonID": "22708898-4460-4b52-9f05-477cbe2a904e"
}

```
Response:
```
{
    "Salon": {
        "Rows": [
            {
                "Blocks": [
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "1A",
                                "SeatClass": "window"
                            },
                            {
                                "Assigned": false,
                                "Num": "1B",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "1C",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "1D",
                                "SeatClass": "middle"
                            },
                            {
                                "Assigned": false,
                                "Num": "1E",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "1F",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "1G",
                                "SeatClass": "window"
                            }
                        ]
                    }
                ]
            },
            {
                "Blocks": [
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "2A",
                                "SeatClass": "window"
                            },
                            {
                                "Assigned": false,
                                "Num": "2B",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "2C",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "2D",
                                "SeatClass": "middle"
                            },
                            {
                                "Assigned": false,
                                "Num": "2E",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "2F",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "2G",
                                "SeatClass": "window"
                            }
                        ]
                    }
                ]
            },
            {
                "Blocks": [
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "3A",
                                "SeatClass": "window"
                            },
                            {
                                "Assigned": false,
                                "Num": "3B",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "3C",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "3D",
                                "SeatClass": "middle"
                            },
                            {
                                "Assigned": false,
                                "Num": "3E",
                                "SeatClass": "aisle"
                            }
                        ]
                    },
                    {
                        "Seats": [
                            {
                                "Assigned": false,
                                "Num": "3F",
                                "SeatClass": "aisle"
                            },
                            {
                                "Assigned": false,
                                "Num": "3G",
                                "SeatClass": "window"
                            }
                        ]
                    }
                ]
            }
        ]
    }
}
```

### AssignSeat assigns seat on given salon ID according to assignation order
Request: 
```
localhost:8080/seat/assign
{
	"SalonID": "22708898-4460-4b52-9f05-477cbe2a904e"
}

```
Response:
```
{
    "AssignedSeat": {
        "Assigned": true,
        "Num": "1B",
        "SeatClass": "aisle"
    }
}
```

### SeatInfo returns information about seat on given index
Request: 
```
localhost:8080/seat/info
{
	"SalonID": "22708898-4460-4b52-9f05-477cbe2a904e",
	"SeatNumber": 500
}

```
Response:
```
{
    "AssignedSeat": {
        "Assigned": false,
        "Num": "1D",
        "SeatClass": "window"
    }
}
```
