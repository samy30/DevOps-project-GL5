db.auth('root','root')

db = db.getSiblingDB('devopsProjectDB')

db.createUser({
	user: 'user',
	pwd: 'user',
	roles: [
	  {
	    role: 'dbOwner',
	    db: 'devopsProjectDB'
	  }
	]
});
