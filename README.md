### MongoDb

#### Enable Text Search
    mongod --setParameter textSearchEnabled=true

#### Create Text Index
    db.ad.dropIndexes()
    db.ad.ensureIndex( { description: "text", title: "text" } );

    db.ad.runCommand("text", { search: "hsjsjd", limit: 20, project: { "price" : 1, "image1": 1 }})
    {
    	"queryDebugString" : "hsjsjd||||||",
    	"language" : "english",
    	"results" : [
    	{
    		"score" : 1.1,
    		"obj" : {
    			"_id" : ObjectId("53202376058424b87c9d9368"),
    			"price" : NumberLong(6468),
    			"image1" : {
    				"egg" : "0001_c0448ef44bf7fd00476fadf11805d94fe94a5820_655B4C_816_612",
    				"baby" : "0001_68415c85528ccf9e763eb48d9dc0fca8a540f701_655B4C_400_300",
    				"infant" : "0001_9b0e36f28be91dae81a02863fadce2bc2f196312_655B4C_200_150",
    				"newborn" : "0001_72a53f664db6f415e9e862c607d9c0ba177c20af_655B4C_100_75"
    			}
    		}
    	},
    	...
    	],
    	"stats" : {
    		"nscanned" : 2,
    		"nscannedObjects" : 0,
    		"n" : 2,
    		"nfound" : 2,
    		"timeMicros" : 96
    	},
    	"ok" : 1
    }