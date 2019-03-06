# datastore-keyloader-demo

> Custom datastore.Keyloader of Datastore Golang SDK.

If one property have mutil types in one Datastore entity, when get entity, it will cause ErrFieldMismatch error.  
Customizing Entity Keyloader can slove this problem.