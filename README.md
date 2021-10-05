# go-marvel-api
using go search the Marvel universe characters via marvel api
</br>
</br>

## Build and run tests on the local environemnt

</br>

### Build the project

> $ go build -a -v -o out/marvel .

</br>

### Run tests

</br>

* Run unit tests
    > $ go test ./... -v -coverprofile="out/test-reports/cover.out"

* Generate unit test  report
    > $ gotestsum --junitfile out/test-reports/unit-tests-report.xml

* Generate coverage report
    > $ go tool cover -html="out/test-reports/cover.out" -o="out/test-reports/cover-report.html"
    
    ![cover-report](assets\cover-report.png)

* Generate func based coverage report
    > $ go tool cover -func="out/test-reports/cover.out" > out/test-reports/cover-func-report.out

    ![cover-func-based](assets\cover-func-based.png)

</br>
</br>

## Circleci integration

</br>

You can join the Circleci Project Team <a href="https://app.circleci.com/pipelines/github/koseburak/go-marvel-api?invite=true" target="_blank">here...</a>

</br>
</br>

[build] stage run after commit pushed to remote for all branches in Circleci workflows.
![circleci-build-release](assets\circleci-build-release.png)
</br>

You can find the build artifact and test reports under Circleci Artifacts menu of current build.
![circleci-build-artifacts](assets\circleci-build-artifacts.png)
</br>

* [publish-github-release] stage is running only after the tag pushed to remote. It's publish the new release automatically including the binary/assets with your tag name on Github.
![github-auto-releae-output](assets\github-auto-releae-output.png)
</br>
