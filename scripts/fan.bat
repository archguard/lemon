echo off

echo Welcome use fan in/out tools
set codePath=%cd%
set /p codePath=Please enter code absolute path (default is current path):

salt_cli fan -s=%codePath% -c=fan.json

lemon_cli ag -s=fan-in.json -c=a.json
echo End, please view result: bad_smell_report.json,suggestion.json