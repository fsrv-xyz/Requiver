![Logo](https://agile-defense.com/wp-content/uploads/2021/06/CPaaS-01.png)


# Requive

A web server written in GoLang, which is used to confirm http requests.


## Run Locally

Clone the project

```bash
git clone https://github.com/enforcer-GH/Requiver.git
```

switch into the Project Directory

```bash
cd Requiver/
```

Install dependencies

```bash
sudo apt install golang
```

Start the server

```bash
./requive 
```


## Parameter

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-web.listen-address` | `string` | **Default**: :8080 |

## Endpoints

| Endpoint | Description                |
| :-------- | :------------------------- |
| `/ping` | used to add client IP Address |
| `/status` | shows all stored IP Addresses |
| `/ack/<specific_ip>` | removes an specific IP Address |
| `/flush` | removes all stored IP Addresses |
