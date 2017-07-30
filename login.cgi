# login.py
import cgi
import cgitb

cgitb.enable()

form = cgi.FieldStorage()

if "acct" not in form or "ppass" not in form:
    print "<H1>Error</H1>"
    print "Please fill in the account and password fields."
    return

print "<p>name:", form["name"].value
print "<p>addr:", form["addr"].value

acct = form.getvalue('acct')
pw = form.getvalue('pass')

if acct == 'randy' and pw == 'fuckingdiaz':
    
