FROM tomcat:8.5-alpine
MAINTAINER BIT  
ADD bonusPointsExchange/bonusPointsExchange.war /usr/local/tomcat/webapps
CMD ["catalina.sh", "run"]
