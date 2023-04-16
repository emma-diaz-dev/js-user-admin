db.createUser({
    user: 'root',
    pwd: 'toor',
    roles: [
        {
            role: 'readWrite',
            db: 'testDB',
        },
    ],
});

db = new Mongo().getDB("testDB");

db.createCollection('user', { capped: false });

db.users.insert([
    { "name": "Emma123" , "last_name":"Diaz","address":"Pj. Durrieu 5432, Cordona, Argentina"},
    { "name": "Juan456" , "last_name":"Pere","address":"San Martin 2435, Cordona, Argentina"},
    { "name": "Jose789" , "last_name":"Diaz","address":"Santa Ana 2453, Cordona, Argentina"},
]);