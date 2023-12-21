db = db.getSiblingDB('go-boilerplate');
db.createCollection("run_time")

db.run_time.insert({
    started_at: new Date()
})