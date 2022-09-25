from twilio.rest import Client 
 
account_sid = 'AC7c1d4068211dfa361cfc6be3a3af78a8' 
auth_token = '[AuthToken]' 
client = Client(account_sid, auth_token) 
 
message = client.messages.create(  
                              messaging_service_sid='MG3f7e55eb345b78cdffe4b6d38f114167', 
                              body='This is the test to see how the twilio API works ...',      
                              to='+16475130152' 
                          ) 
