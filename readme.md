## Usage : go run main.go <path file> <flag> <detail>
### Option :
> -t                     Flag For Option to Convert the File to Plaintext (text) or JSON (json)
> 
> -o                     Flag For Choose Where He Will Put the Output File. [/home/$USER/text.txt], [/home/$USER/text.json]
> 
> -h                     Flag For Instructions for Use.
>

### List File :
> /var/log/apache2/error.log
> 
> /var/log/dpkg.log
> 
> /var/log/kern.log
> 
> /var/log/auth.log

### Example :
> go run main.go /var/log/auth.log                                       Output File txt
> 
> go run main.go /var/log/auth.log -t text                               Output File txt
> 
> go run main.go /var/log/auth.log -t json                               Output File json
> 
> go run main.go /var/log/auth.log -t json -o /home/$USER/text.json      Output File in Folder /home/$USER/ with Type Json and File Name text.json
> 
> go run main.go /var/log/auth.log -o /home/$USER/text.txt               Output File In Folder /home/$USER/ with File Name text.txt
> 
> go run main.go -h  