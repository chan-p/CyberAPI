echo "====API TEST START===="

echo "TEST:show:M2"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M2
echo ""
echo "Correct Answer"
echo "{"name":["犬塚　健太","後藤　紳","佐々木　崇"]}"
echo ""

echo "ERROR_TEST:show:M3"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M3
echo ""
echo "Correct Answer"
echo "{"massege":"not exit grade"}"
echo ""

echo "TEST:show:team_count:RC"
echo "Your Answer"
curl http://localhost:50000/api/show/team_member_count?team=RC
echo ""
echo "Correct Answer"
echo "{"projectmembercount":3}"
echo ""

echo "ERROR_TEST:show:team_count:TEST"
echo "Your Answer"
curl http://localhost:50000/api/show/team_member_count?team=TEST
echo ""
echo "Correct Answer"
echo "{"massege":"not exit team"}"
echo ""

echo "TEST:show:team_count:CA"
echo "Your Answer"
curl http://localhost:50000/api/show/team_member_count?team=CA
echo ""
echo "{"projectmembercount":2}"
echo ""

echo "TEST:show:M1"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M1
echo ""
echo "Correct Answer"
echo "{"name":["宗政 一舟","友松　祐太","林　知範","譚　敬元"]}"
echo ""

echo "TEST:data:member_add:New_Member,M1,test@gmail.com,RC"
echo "Your Answer"
curl http://localhost:50000/api/data/member_add?name=New_Member\&grade=M1\&mail_address=test88@gmail.com\&team=RC
echo ""
echo "Correct Answer"
echo "{"massege":"OK"}"
echo ""

echo "TEST:show:M1"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M1
echo ""
echo "{"name":["宗政 一舟","友松　祐太","林　知範","譚　敬元","New_Member"]}"
echo ""

echo "TEST:show:M2"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M2
echo ""
echo "Correct Answer"
echo "{"name":["犬塚　健太","後藤　紳","佐々木　崇"]}"
echo ""

echo "TEST:show:team_count:RC"
echo "Your Answer"
curl http://localhost:50000/api/show/team_member_count?team=RC
echo ""
echo "Correct Answer"
echo "{"projectmembercount":4}"
echo ""

echo "TEST:data:member_update:New_Member,M2"
echo "Your Answer"
curl http://localhost:50000/api/data/member_update?name=New_Member\&new_grade=M2\&new_team=CL
echo ""
echo "Correct Answer"
echo "{"massege":"OK"}"
echo ""

echo "TEST:show:team_count:CL"
echo "Your Answer"
curl http://localhost:50000/api/show/team_member_count?team=CL
echo ""
echo "Correct Answer"
echo "{"projectmembercount":4}"
echo ""

echo "TEST:show:M1"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M1
echo ""
echo "Correct Answer"
echo "{"name":["宗政 一舟","友松　祐太","林　知範","譚　敬元"]}"
echo ""

echo "TEST:show:M2"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M2
echo ""
echo "Correct Answer"
echo "{"name":["犬塚　健太","後藤　紳","佐々木　崇","New_Member"]}"
echo ""

echo "TEST:data:delete:New_Member"
echo "Your Answer"
curl http://localhost:50000/api/data/member_delete?name=New_Member
echo ""
echo "Correct Answer"
echo "{"massege":"OK"}"
echo ""

echo "TEST:show:M2"
echo "Your Answer"
curl http://localhost:50000/api/show/grade_all?grade=M2
echo ""
echo "Correct Answer"
echo {"name":["犬塚　健太","後藤　紳","佐々木　崇"]}

echo "====API TEST END===="
