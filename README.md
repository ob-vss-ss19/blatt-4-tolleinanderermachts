[![Build Status](https://travis-ci.com/ob-vss-ss19/blatt-4-tolleinanderermachts.svg?branch=develop)](https://travis-ci.com/ob-vss-ss19/blatt-4-tolleinanderermachts)

Contributors:
---
* Marcel Reineck
* Andreas Stiglmeier
 
---
 
How to run the services:
---
##### Download docker images:
````
terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-4-tolleinanderermachts:feature-docker-moviecontrol
````
````
terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-4-tolleinanderermachts:feature-docker-roomcontrol
````
````
terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-4-tolleinanderermachts:feature-docker-showcontrol
````
terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-4-tolleinanderermachts:feature-docker-usercontrol

terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-4-tolleinanderermachts:feature-docker-reservationcontrol

##### Alternative: Run the 5 services in the 5 "*control" subfolders.

---

##### Fill Client with Dummy Data:
````$xslt
go run ./client/dataclient.go
````

##### Please Note:
Due to the services communicating, the port range from 50.000 to 70.000 is exposed on the docker images.
````
2019/06/24 12:37:19 Transport [http] Listening on [::]:52682
2019/06/24 12:37:19 Broker [http] Connected to [::]:52683
````

Communication
---
For further information, see [protocol.md](https://github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/blob/develop/protocol.md)
 
