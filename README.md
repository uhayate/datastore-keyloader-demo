# datastore-keyloader-demo

> Custom Keyloader of Datastore Golang SDK.

If one property have mutil type in one Datastore entity, when get entity, it will cause ErrFieldMismatch error.  
Custom Entity Keyloader can slove this problem.