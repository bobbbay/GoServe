binary=`uname -m`

git clone https://github.com/Bobbbay/GoServe.git && curl --output ./GoServe/GoServe "https://raw.githubusercontent.com/Bobbbay/GoServe/master/binaries/GoServe-${binary}" && cd ./GoServe && chmod u+x ./GoServe && rm -rf makefile && wget https://raw.githubusercontent.com/Bobbbay/GoServe/master/binaries/makefile && cd .. && echo "\n # Added by GoServe\n" export PATH=$PWD + $PATH >> .bashrc
