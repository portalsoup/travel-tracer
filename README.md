# travel-tracer

strava-like web app for tracking geographical and personal metrics while performing outdoor activities.  This is
more a learning sandbox for me than a repo where anything useful will emerge.

This app requires a MongoDB connection that uses no credentials.  
This can be set with the command line arg db.host.  Example: -db.host=127.0.0.1

Current endpoints:
    
   GET /map/point 
   GET /map/route
   GET /coordinates/:coordinateId
   GET /coordinates
   GET /route
   PUT /db/saveCoordinate ? id=coordinateId