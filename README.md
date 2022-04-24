# Cyclic start and command commands in the terminal

[Github](https://github.com/SidorkinAlex/bashCommandCyclicRunner) |
[Ru](#Циклический-запускикоманд-команд-в-терминале) |
[En](#Cyclic-start-and-command-commands-in-the-terminal)

![image](http://web-seedteam.ru/wp-content/uploads/2021/04/screenshot-0.0.0.0-2021.04.15-23_18_20.png)
![image](http://web-seedteam.ru/wp-content/uploads/2021/04/screenshot-0.0.0.0-2021.04.15-23_43_05.png)

The program is designed to run commands at a fixed time interval (specified in seconds)

## Installation

Clone the repository of the latest stable version and run the build of the project with the command
```
make
```
Or download the compiled binary from the link

The Latest Release:



## Usage

### Creating A Task

create a command file.json in the directory where the executable file is located

```
[
{
"interval": 10,
"command": "php 1.php"
},
{
"interval": 1,
"command": "python script.py"
},
{
"interval": 30,
"command": "ls -l"
}
]
```
interval - the time in seconds after which the command
command will be run - the command that will be executed at the above interval

or run the command:
```
./main -create-job

```

After completing all the steps, a cyclic command will be added

important:

To apply a new command, a program restart is required

### Start


To run, while in the directory with the file, run the command

```
./main
```
after the information message appears

2022/04/19 22:09:27 Programm bashCommandCyclicRunner has been success running

### Stop

to stop the execution of cyclic commands, run the command

```
./main -stop
```

### Restart

To restart, run the following command:

```
./main -restart
```

### License
MIT


# Циклический запускикоманд команд-в терминале

[Github](https://github.com/SidorkinAlex/bashCommandCyclicRunner) |
[Ru](#Вебхуки-для-SuiteCRM) |
[En](#Webhooks-from-SuiteCRM)

[Установка](#Установка)

Программа предназначена для запуска команд через фиксированный интервал времени (указывается в секундах)

## Установка

Клонируйте репозиторий последней стабильной версии и запустите сборку проекта командой
```
make
```
Или скачайте скомпилированный бинарник по ссылке

Последний Релиз:



## Использование

### Создание Задания

Создайте файл command.json в директории в которой лежит исполняемый файл

   ```
   [
  {
    "interval": 10,
    "command": "php 1.php"
  },
  {
    "interval": 1,
    "command": "python script.py"
  },
  {
    "interval": 30,
    "command": "ls -l"
  }
]
   ```
interval - время в секундах, через которое будут запущена команда.
command - команда, которая будет выполняться с вышеуказанным интервалом

Создание Задания из приложения - выполните команду:
```
./main -create-job

```

Пройдя все этапы будет добавлена циклическая команда

ВАЖНО:

Для применения новой команды требуется рестарт программы

### Запуск


Для запуска, находясь в дирректории с файлом выполните команду

   ```
 ./main
   ```
После появления информационного сообщения

2022/04/19 22:09:27 Programm bashCommandCyclicRunner has been success running

### Остановка

для остановки выполнения циклических команд выполните команду

```
./main -stop
```

### Рестарт

Для перезапуска выполните следующую команду:

```
./main -restart
```


### License
MIT