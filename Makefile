
# Start Atlas dependencies using Docker Compose.
env-up:
	@echo "Starting atlas dependencies..."
	docker-compose -f deployment/docker_compose.yaml up -d
	make rs-initiate

# Initialize a MongoDB replica set.
rs-initiate:
	@echo "Initializing mongo replica set..."
	docker exec -it mongo1 mongosh --host mongo1:4010 --eval "rs.initiate({_id:'athena0', members: [{_id:0, host: 'mongo1:4010', "priority": 1},{_id:1, host: 'mongo2:4020', "priority": 2},{_id:2, host: 'mongo3:4030', "priority": 3}]}, {force: true})"