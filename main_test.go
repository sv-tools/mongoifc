package mongoifc_test

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

const (
	DockerImage   = "mongoifc"
	DockerFile    = ".github/test.dockerfile"
	DockerName    = "mongoifc-test"
	MongoPort     = "27888"
	MongoUsername = "admin"
	MongoPassword = "adminpass"
	MongoUri      = "mongodb://" + MongoUsername + ":" + MongoPassword + "@127.0.0.1:" + MongoPort +
		"/?authSource=admin&directConnection=true"
)

func TestMain(m *testing.M) {
	fmt.Println("Building docker image...")
	cmd := exec.Command(
		"docker",
		"build", ".",
		"--rm",
		"--file", DockerFile,
		"--tag", DockerImage,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("docker build failed:\n%s", string(out))
		panic(err)
	}

	fmt.Println("Starting docker container...")
	cmd = exec.Command(
		"docker",
		"run",
		"-d",
		"--name="+DockerName,
		"-p", MongoPort+":27017",
		"-e", "MONGO_INITDB_ROOT_USERNAME="+MongoUsername,
		"-e", "MONGO_INITDB_ROOT_PASSWORD="+MongoPassword,
		DockerImage,
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("docker run failed:\n%s", string(out))
		panic(err)
	}
	defer func() {
		fmt.Println("Stopping docker container...")
		cmd := exec.Command(
			"docker",
			"rm",
			"--force",
			DockerName,
		)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("docker rm failed:\n%s", string(out))
			panic(err)
		}
	}()

	fmt.Println("Waiting docker container...")
	time.Sleep(10 * time.Second)

	fmt.Println("Initializing mongodb...")
	cmd = exec.Command(
		"docker",
		"exec", DockerName,
		"/usr/bin/mongosh",
		"-u", MongoUsername,
		"-p", MongoPassword,
		"--eval", "rs.initiate()",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("docker exec failed:\n%s", string(out))
		panic(err)
	}

	fmt.Println("Running tests...")
	if code := m.Run(); code != 0 {
		panic(fmt.Sprintf("tests failed with code %d", code))
	}
}
