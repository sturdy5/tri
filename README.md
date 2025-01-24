# Tri

This is a todo tutorial I put together as part of this workshop - <https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/>.

## Usage

You will first need to define the file that the todos will be stored in. There are two ways that you can do this.

1. You can set an environment variable with the path to the json file

    ```shell
    export TRI_DATAFILE=~/.todos.json
    ```

1. You can create a yaml file in your home directory called `.tri.yaml` that specifies the path

    ```yaml
    datafile: ~/.todos.json
    ```

Either of those configurations can work. Now you can run the command

```shell
./tri --help
```

That will print out usage statement for the command. You can pass in the `--help` flag into all the subcommands as well to learn more about how they work.
