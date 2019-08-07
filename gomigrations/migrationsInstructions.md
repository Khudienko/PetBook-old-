1. Installation
   To install the library and command line program, use the following:
   
   go get -v github.com/rubenv/sql-migrate/...
2. Available commands are:

       down      Undo a database migration
       new       Create a new migration
       redo      Reapply the last migration
       status    Show migration status
       up        Migrates the database to the most recent version available
3. Specify the parameters of your database in the dbconfig.yml file 
4. Run the command 'sql-migrate up' 