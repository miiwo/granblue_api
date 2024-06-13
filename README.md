# Skyfarer REST API App
![](https://img.shields.io/badge/public_API_\@api.skyfaring--domain\.xyz-online!-brightgreen)

[/ For red: D24939 | For green: brightgreen /]: #

![-banner picture here-](SkyfarerAPIBanner.png)

## Introduction
This API grabs item info for the game "Granblue Fantasy" in a RESTful manner. This will allow for easier searching of certain properties on items and their magnitude. Planned usage to use this in conjecture with a website to search the collection of items. This app is written in Golang.

For usage, go [here](#usage).

## Features
- Authentication via an API key
- Find GBF weapons with filters based on weapon skills to help with grid building
- Publically accessable REST API that can be used by other developers without needing to source their own data

## Setup on your own machine
#### 1. Grab the source code
- Clone the repository using this URL: `https://github.com/miiwo/granblue_api.git`

#### 2. Setup the Database 

- Ensure you have a MySQL-flavored DB server installed.

[Setup the database by running: `setup.sql` in the DB server]: #

- If you are running the database server on a different host machine than this app, please look into allowing remote access to the DB server.

#### 3. Setup the Application

- Make sure you have `go` (version 1.22.3+) installed on your machine. *(Check by typing `go version` into the Command Terminal)*  
- Run `go build` in the Command Terminal to grab all the dependencies and to build the app.  
- Make a `.env` with the following variables. See `.env.sample` for assistance.

| Key       | Notes                                                                                 |
| ---       | ---                                                                                   |
| API_KEY   | Used to access the API                                                                |
| DB_USER   | Name of DB user to access the API                                                     |
| DB_PASS   | Password of associated DB user to access the API                                      |
| DB_HOST   | The host address of the DB server                                                     |
| DB_PORT   | Port of the DB server                                                                 |
| DB_NAME   | Name of the database to use                                                           |
| BASE_URL  | URL used to access the API. Must be in the form: `<host address>:<port>` or `:<port>` |


## Usage
This API is publically hosted at: `api.skyfaring-domain.xyz`  
| Is it Up? |
| :-:       |
| :white_check_mark:       | 

[/:white_check_mark: vs. :x: /]: #

You can connect via Postman at the endpoint: `https://api.skyfaring-domain.xyz/v1/weapons`

Or use cURL on the Command Terminal:  
Ex. `curl -H "Authorization: Bearer <API_KEY>" https://api.skyfaring-domain.xyz/v1/weapons?name=phoenix `

> [!NOTE]
> I try to keep it up during 9am-5pm EST hours, but it may be down due to cloud hosting cost concerns :(  
> For API Key access, please make an issue on the repository requesting access.  
---

To self-host the app, run `go run .` to start the application.

Please fill in the Authorization with a Bearer Token containing the <*API_KEY*> and use the following endpoints[ here](#endpoints)!

Connect to the API using Postman on the port and <*API_KEY*> you set in the `.env` file.  
You can connect via Postman at the endpoint: `http://<BASEURL>/v1/weapons`

Or use cURL on the Command Terminal:  
Ex. `curl -H "Authorization: Bearer <API_KEY>" http://<BASEURL>/v1/weapons?name=phoenix `

## Endpoints

| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /v1/weapons | To get information on a set of weapons                |
| GET | /v1/weapons/:name | To get information on a particular weapon       |
| GET | /ping       | Gives you pong back without needing to authenticate   |

[| GET | /v1/characters | To get information on a character |]: #

#### `GET v1/weapons`
This endpoint will grab a list of weapons from the database. Returns at max 20 entries.

##### Request
`GET http://<BASEURL>/v1/weapons`  
Fields to query by:
- name
- element
- wep_type

##### Response
Returns a list of weapons. If there are no skills associated with the weapon, there will be no field for it in the response JSON.
```
[
    {
        "Name": "Abyss Gaze",
        "Element": "dark",
        "WeaponType": "gun",
        "Description": "Within the void, darkness reigns supreme. A torrent of cataclysmic power manifests as a gun, casting a shadow across the skies. Where there is light there is shadow, turning to darkness and dissolving into nothingness.",
        "OugiName": "Quadra",
        "OugiDesc": "Massive Dark damage to a foe. Inflict Blinded and 10% Debuff Res. Down",
        "LvlOneAtk": 424,
        "LvlHundredAtk": 2518,
        "LvlOnefiftyAtk": 2940,
        "LvlTwohundredAtk": 0,
        "Skills": [
            {
                "Name": "Mistfall's Might",
                "Description": "Medium boost to dark allies' ATK",
                "BoostType": "MAGNA",
                "StatAffected": "ATK",
                "SkillLvlOne": "3.0",
                "SkillLvlTen": "12.0",
                "SkillLvlFifteen": "14.5",
                "SkillLvlTwenty": ""
            },
            {
                "Name": "Mistfall's Restraint",
                "Description": "Medium boost to dark allies' double attack rate and critical hit rate",
                "BoostType": "MAGNA",
                "StatAffected": "DA/CRIT",
                "SkillLvlOne": "0.8/4.4",
                "SkillLvlTen": "3.5/8.0",
                "SkillLvlFifteen": "5.0/10.0",
                "SkillLvlTwenty": ""
            }
        ]
    },
    ...
]
```

## Technologies Used
- [Golang]() | The programming language used to write the REST API app. Also known as Go
- [Gin]() | A Golang web framework featuring high performance and productivity
- [GORM]() | A ORM framework that simplifies modelling the database schema within the development of the app
- [MariaDB]() | A free open-source SQL database
- [Google Cloud Platform]() | A cloud provider used to publically host the REST API app. Specifically using Compute Engine which is not serverless

## License
This project is available for use under the MIT License.
