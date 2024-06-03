/* This code is part of RF Switch by @Penthertz
*  Author(s): Sébastien Dudek (@FlUxIuS)
*/

package main

import (
	"fmt"

	cli "penthertz/rfswift/cli"
)


var version = "0.3"

var ascii_art = `                                                                                                    
                                                                                          
                                                                                          
                      :===:                                                               
                   .*@*=:-*%@=.                     .:*@@@@@@@@@@%+:                      
                 :*@*+*=:::-:=@*.                 -%@%=**+==-:::::-%@%:                   
               =@*==%=-:  ..  .=%-              =@@=-*+===-::.  ::::::+%-                 
              :%.*::*@=-.   .+::.@+           :@@=:*+===-  .=**+==-::-*%+%.               
            .#@*+=...:*+*+...::*=:*%.        +@%.+=+==:.:=****+-:.  .:::.-#:              
           .#***%%+-:+=  .=*.:.=#%-=%:     .%@=.*+..:.:==***=..::::...  .=*@=.            
          .+.%%*%%=-*: :+: .*=::.:%%=%+   .@@:=%*=-: -==..- :=-::.    :-:   *%=           
          :.=@@**#**:-%@@#%* -*-:*-.:+#@%#@%=*%#%= :==+*+:.-=:.  :-==:.:=-=::*+#.         
          .:@@=%++*:  :*++*:  .#%==@@#=-:::-+##:  :===++:.--.:#@@@@@@@@@==%==#=*@.        
           #@#%#=*=          -+@%#*%*%@@@@#=: : :====++: :.:%@%*%@@%%%*=*@@@**%:@*        
          .%+*=%*=:        .:- :+@=@@@@@@%%#+:=++=-++-  ..#@*    *@*#%*:::.:*@@-*%.       
           %-%-%-%@:        .   *%@*@@*=+****=::=++*#.  =@@:     +@*=:.*@=    =@@%:       
           -*:%@=+@#+    .:::  :%@%%@@@@@@@@@%%@@@+:::=%@*.    :=#@+:+ :=%%::. :%@:       
            :: *@=%#=#*:   . .*@@@@@@*-+#%*%@@%+====+%@%-:=*%%*.-@%=+* -=.*@:+   %%       
             . -@%@#-==#@+.=%@@@@%**@%%@*::===*@@@@@@@@@%*=:. .*@%=+@::=::=@%-*  =@=      
                *%@% :===*%@@@@@@@@%*=*%*=+=::::::=*%@@@@%%%@@@%=-%%..=::=-%%:%. -@+      
                 =@%   :===%@@@#+*@@@@@*:+@#=====:    .:------=*@*  :..=-:=%%=%  =@=      
                 -@%.     :**@@@@%*==*@@@@@%%%%*======-::.         :-=-:=.=@**- .%@.      
             .:=%%%*.     :%%*:.:-=#@@=.:*%@@@@@@@%+::-============-::= :%@*-= .%@:       
          =@@#==%%==    :==%@=       :@%=-..:==*%@@@@@@@%+:.        :+%@@#:+: -@@-        
         %%.  .%@. :      :%@+        .#@%*+=-:::-===+**#%%@@@@@@@@%%#*=: . =%@#.         
        :@+  :@%.         .%@*         .@@%@@@*===-:..:-============:  ..-%@%=..=%%+      
         %%.:%%:           *@%.         .%# .-*@@@%%*=--------:::::-+#%@@*:-*%*-:.+%:     
         .@*@*             :@@- :=.      :@#.     :+#%%@@@@#==+#@@@%#+:.=%#=     -@#      
          *@:     :.        *@@-**=.      .@*             :%@%@#.    =@*:       *@@.      
        :%*.    .=.       .-=%@%==#+:      :@*           :%=:@@##%#%@%#####.  -%@@:       
       .%%.    :*:     :**%@@%@@#==%+.      :@+        ..:%+@@@@@%.     .@: :%@@@-        
       *@@%.  .+*.   *%@@@@@@@@@@#==@+:      :@*  .=#%@@@:**@@@@@+     .*:-%@@@#.         
       %@@@%: =*+: =@@@@@@@%%%@%@@**@@*:     :=*@@%*=:   :%%=***@%:    =%@@@@%@=          
       %@@%@@+=+#=%@@@%%%%%%@@@@@@@@@@@=..*@@%-    .=:.%#%:     #@@@:-@@@@@@=+%.          
       :@@@@@@#=*%@@%%%%@@@@@@@@@@%=::*#*=:  :=*%@*:=%+-%:      :%%@@@@@@@= -%.           
        =@@@@@@@#@@@@@@@@@@@@@%=..-*-   .-#@@%+.  =%= .%*++=====%@@@@@@%:  *%:            
          +@@@@@@@@@@@@@@@%-.       :*%@@*:.    -@=        =%@@@@@@@@@@= :%=              
            *@@@@@@@@@%+:  ..  .=*@@@*-        -%:  .:=*%@@@@@@@%*@@@@%#%*.               
            :@@@@@@%: .:=:. :%@@@%-             #@@@@@@@@@@%+.   .*@@@@%-                 
            .%@@+:.:=-:::*%@@@+::::::..                  :-=+****==#@@#                   
             =@=-=-:-=#@@@%*          .:=+#%%@@@#*=::.          .=.*@@@::.                
             .@*:==%@@@@+.                      ...:=*%@@#=.      .@@@@+=@%:              
              *@#@@@@*:                                    .:==:: =@@@@%. .=#+.           
              .@@@%-                                                          .=          
               **:                                                                                   

			888~-_   888~~        ,d88~~\                ,e,   88~\   d8   
			888   \  888___       8888    Y88b    e    /  "  _888__ _d88__ 
			888    | 888          'Y88b    Y88b  d8b  /  888  888    888   
			888   /  888           'Y88b,   Y888/Y88b/   888  888    888   
			888_-~   888             8888    Y8/  Y8/    888  888    888   
			888 ~-_  888          \__88P'     Y    Y     888  888    "88_/        

						RF toolbox for HAMs and professionals                                                                             
`

func main() {
    fmt.Println(ascii_art)
    fmt.Print("Version: ", version, "\n\n")
    cli.Execute()
}
