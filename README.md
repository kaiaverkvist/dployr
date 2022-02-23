# dployr

Dployr is a tool that extracts expected environment variables from a docker-compose file and makes it simple
to deploy on a remote server.

*This is a screenshot of what you're presented with once you launch dployr against a directory containing a docker-compose.yml.*
![screenshot from dployr](https://i.imgur.com/Isv6e67.png)

## ⚔️ Usage
In the screenshot above, the command used was:

`dployr -d /Documents/Development/personal-site/ --host 192.168.0.97 -u ctl`

This detects the `docker-compose.yml` file, and extracts the environment variables inside the file.
Dployr then proceeds to open your browser and displays what is seen in the screenshot above.


## 🔥 Get the tool
Currently there is no binary distribution available, but this is planned to be available soon.

## 🏗️ Contributions
Contributions are welcome and wanted!

## 🏁 Long term goals
* Deployment to kubernetes clusters, AWS, etc using this relatively simple model.
* ?