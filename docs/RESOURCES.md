# Hue resources

Core concepts Philips Hue API: https://www.developers.meethue.com/documentation/core-concepts

Clip API:
https://developers.meethue.com/documentation/lights-api

## Internal clip debugger
http://192.168.2.2/debug/clip.html  

## Creating a user

Post: {"devicetype":"my_hue_app#iphone peter"}  
To: http://192.168.2.2/api

Resulting in:  
```JSON
[
    {
        "success": {
            "username": "adasdasdsdaasddsdasdsdsacB1249823ujfew"
        }
    }
]

```
