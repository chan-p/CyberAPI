echo "TEST:address_count"
curl http://localhost:1323/api/show/address_count
echo ""
echo "TEST:show:M2"
curl http://localhost:1323/api/show/grade_all?name=M2
echo ""
echo "ERROR_TEST:show:M3"
curl http://localhost:1323/api/show/grade_all?name=M3
echo ""
echo "TEST:show:team_count:AD ANSEWR:2"
curl http://localhost:1323/api/show/team_member_count?name=AD
echo ""
echo "ERROR_TEST:show:team_count:TEST ANSWER:ERROR"
curl http://localhost:1323/api/show/team_member_count?name=TEST
echo ""
echo "TEST:show:team_count:RC ANSWER:3"
curl http://localhost:1323/api/show/team_member_count?name=RC
echo ""
echo "TEST:show:M1 ANSEWR:4"
curl http://localhost:1323/api/show/grade_all?name=M1
echo ""
echo "TEST:data:member_add:New_Member,M1,test@gmail.com,RC"
curl http://localhost:1323/api/data/member_add?name=New_Member&grade=M1&mail_address=test88@gmail.com&project=RC
echo ""
echo ""
echo ""
echo "TEST:show:M1 ANSWER:5"
curl http://localhost:1323/api/show/grade_all?name=M1
echo ""
echo "TEST:show:M2 ANSWER:3"
curl http://localhost:1323/api/show/grade_all?name=M2
echo ""
echo "TEST:show:team_count:RC ANSWER:4"
curl http://localhost:1323/api/show/team_member_count?name=RC
echo ""
echo "TEST:data:member_update:New_Member,M2"
curl http://localhost:1323/api/data/member_update?name=New_Member&new_grade=M2
echo ""
echo "TEST:show:M1 ANSWER:4"
curl http://localhost:1323/api/show/grade_all?name=M1
echo ""
echo "TEST:show:M2 ANSWER:4"
curl http://localhost:1323/api/show/grade_all?name=M2
echo ""
echo "TEST:data:delete:New_Member"
curl http://localhost:1323/api/data/member_delete?name=New_Member
echo ""
echo "TEST:show:M2 ANSWER:3"
curl http://localhost:1323/api/show/grade_all?name=M2
echo ""
