function DockerBuild {

    
    $image_path = "$env:DOCKER_USERNAME/graphql:latest"
    docker login -u $env:DOCKER_USERNAME -p $env:DOCKER_PASSWORD
    docker build -t $image_path .
    docker push $image_path
}


DockerBuild