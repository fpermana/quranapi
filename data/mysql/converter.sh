#!/usr/bin/bash

if [ ! -f "quran-data.xml" ]
then
	exit 0
fi

#<sura index="1" ayas="7" start="0" name="الفاتحة" tname="Al-Faatiha" ename="The Opening" type="Meccan" order="5" rukus="1" />
#grep "<sura " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); ayas=$(echo "${line}" | awk -F \" '{print $4}'); start=$(echo "${line}" | awk -F \" '{print $6}'); name=$(echo "${line}" | awk -F \" '{print $8}'); tname=$(echo "${line}" | awk -F \" '{print $10}'); ename=$(echo "${line}" | awk -F \" '{print $12}'); type=$(echo "${line}" | awk -F \" '{print $14}'); order=$(echo "${line}" | awk -F \" '{print $16}'); rukus=$(echo "${line}" | awk -F \" '{print $18}'); echo "insert into suras (id,ayas,start,name,tname,ename,type,order,rukus) values ('${id}', '${ayas}', '${start}', '${name}', '${ename}', '${type}', '${order}', '${rukus}');"; done > suras.sql
echo "DROP TABLE IF EXISTS suras;" > suras.sql
echo -n "CREATE TABLE suras (id int(3) NOT NULL PRIMARY KEY, ayas int(3) NOT NULL default 0, start int(4) NOT NULL default 0, name text NOT NULL, tname text NOT NULL, ename text NOT NULL, type text NOT NULL, \`order\` int(3) NOT NULL default 0, rukus int(3) NOT NULL default 0)" >> suras.sql
#echo -n "ENGINE=MyISAM DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci" >> suras.sql
echo ";" >> suras.sql
grep "<sura " quran-data.xml | sed 's|^[ \t]*<sura index="|insert into suras (id,ayas,start,name,tname,ename,type,\`order\`,rukus) values (|g;s|" ayas="|,|g;s|" start="|,|g;s|" name=|,|g;s| tname=|,|g;s| ename=|,|g;s| type=|,|g;s| order="|,|g;s|" rukus="|,|g;s|"[ \t]*/>|);|g;' >> suras.sql

#<juz index="1" sura="1" aya="1" />
#grep "<juz " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); echo "insert into juzs (id,sura,aya) values ('${id}', '${sura}', '${aya}');"; done > juzs.sql
echo "DROP TABLE IF EXISTS juzs;" > juzs.sql
echo "CREATE TABLE juzs (id int(2) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0);" >> juzs.sql
grep "<juz " quran-data.xml | sed 's|^[ \t]*<juz index="|insert into juzs (id,sura,aya) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|"[ \t]*/>[ \t]*|);|g;' >> juzs.sql

#<quarter index="1" sura="1" aya="1" />
#grep "<quarter " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); echo "insert into hibzs (id,sura,aya) values ('${id}', '${sura}', '${aya}');"; done > hibzs.sql
echo "DROP TABLE IF EXISTS hibzs;" > hibzs.sql
echo "CREATE TABLE hibzs (id int(2) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0);" >> hibzs.sql
grep "<quarter " quran-data.xml | sed 's|^[ \t]*<quarter index="|insert into hibzs (id,sura,aya) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|"[ \t]*/>[ \t]*|);|g;' >> hibzs.sql

#<manzil index="1" sura="1" aya="1" />
#grep "<manzil " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); echo "insert into manzils (id,sura,aya) values ('${id}', '${sura}', '${aya}');"; done > manzils.sql
echo "DROP TABLE IF EXISTS manzils;" > manzils.sql
echo "CREATE TABLE manzils (id int(2) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0);" >> manzils.sql
grep "<manzil " quran-data.xml | sed 's|^[ \t]*<manzil index="|insert into manzils (id,sura,aya) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|"[ \t]*/>[ \t]*|);|g;' >> manzils.sql

#<ruku index="1" sura="1" aya="1" />
#grep "<ruku " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); echo "insert into rukus (id,sura,aya) values ('${id}', '${sura}', '${aya}');"; done > rukus.sql
echo "DROP TABLE IF EXISTS rukus;" > rukus.sql
echo "CREATE TABLE rukus (id int(2) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0);" >> rukus.sql
grep "<ruku " quran-data.xml | sed 's|^[ \t]*<ruku index="|insert into rukus (id,sura,aya) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|"[ \t]*/>[ \t]*|);|g;' >> rukus.sql

#<page index="1" sura="1" aya="1" />
#grep "<page " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); echo "insert into pages (id,sura,aya) values ('${id}', '${sura}', '${aya}');"; done > pages.sql
echo "DROP TABLE IF EXISTS pages;" > pages.sql
echo "CREATE TABLE pages (id int(3) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0);" >> pages.sql
grep "<page " quran-data.xml | sed 's|^[ \t]*<page index="|insert into pages (id,sura,aya) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|"[ \t]*/>[ \t]*|);|g;' >> pages.sql

#<sajda index="1" sura="7" aya="206" type="recommended" />
#grep "<page " quran-data.xml | while read line; do id=$(echo "${line}" | awk -F \" '{print $2}'); sura=$(echo "${line}" | awk -F \" '{print $4}'); aya=$(echo "${line}" | awk -F \" '{print $6}'); type=$(echo "${line}" | awk -F \" '{print $8}'); echo "insert into sajdas (id,sura,aya) values ('${id}', '${sura}', '${aya}', '${type}');"; done > sajdas.sql
echo "DROP TABLE IF EXISTS sajdas;" > sajdas.sql
echo "CREATE TABLE sajdas (id int(2) NOT NULL PRIMARY KEY, sura int(3) NOT NULL default 0, aya int(3) NOT NULL default 0, type text NOT NULL);" >> sajdas.sql
grep "<sajda " quran-data.xml | sed 's|^[ \t]*<sajda index="|insert into sajdas (id,sura,aya,type) values (|g;s|" sura="|,|g;s|" aya="|,|g;s|" type=|,|g;s|[ \t]*/>[ \t]*|);|g;' >> sajdas.sql
