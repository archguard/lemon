echo off

echo Welcome use bad smell detect tools.
set codePath=%cd%
set /p codePath=Please enter code absolute path (default is current path):
set /p needInclude=Code is c or cpp?(y/n):
if %needInclude% == y salt_cli include -s=%codePath% -c=code_check_config.json
echo Header file end, begin do bad smell

salt_cli bs -s=%codePath% -c=code_check_config.json
echo End, please view result: bad_smell_report.json,suggestion.json