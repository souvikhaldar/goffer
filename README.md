# Goffer
A powerful tool for fuzzing web applications, in a fast concurent fashion. Written in go.

# Features
1. Fuzzing 
   Fuzzing a program named `OVERFLOW1` server on machine `10.10.85.251` running on port 1337 
   ![image](https://user-images.githubusercontent.com/20991754/121797386-7bfa3a80-cc3d-11eb-865c-e46db0a85e82.png)
You can pass `-s` to increase the number of concurrently running goroutines, by default it is 1. It is advised to keep it low because it depends on the threading capability of the target server.
   
