---
title: SSE protocol
---

ETOS communicates to clients using SSE (server sent events). For now
ETOS only sends log messages using the protocol but we are adding more
data about testruns into the protocol.

  ----------------------------------------------------------------------------------
  Name              Why               Example                      State
  ----------------- ----------------- ---------------------------- -----------------
  Ping              Keep the          {id: None, event: Ping,      Implemented
                    connection alive  data: \"\"}                  

  Shutdown          Server wants the  {id: None, event: Shutdown,  Implemented
                    client to shut    data: \"ESR has shut down\"} 
                    down the                                       
                    connection                                     

  NotFound          The server cannot {id: None, event: NotFound,  Suggested
                    find the ESR      data: \"ESR not found\"}     
                    instance, retry                                
                    after a while                                  

  Error             The server got an {id: None, event: Error,     Suggested
                    error. Retry may  data: \'{\"retry\": bool,    
                    be possible       \"reason\": \"Some sort of   
                                      error\"}}                    

  Message           A user log        {id: 1, event: message,      Implemented
                    message from ETOS data: \"{\'message\': \'a    
                                      message\', \'@timestamp\':   
                                      \'%Y-%m-%dT%H:%M:%S.%fZ\',   
                                      \'levelname\': \'INFO\',     
                                      \'name\': \'ESR\'}\"}        

  Artifact          An artifact to    {id: 1, event: Artifact,     Implemented
                    download          data: \"{\'url\':            
                                      \'<http://download.me>\',    
                                      \'name\': \'filename.txt\',  
                                      \'directory\':               
                                      \'MyTest_SubSuite_0\',       
                                      \'checksums\': {\'SHA-224\': 
                                      \'\<hash\>\', \'SHA-256\':   
                                      \'\<hash\>\', \'SHA-384\':   
                                      \'\<hash\>\', \'SHA-512\':   
                                      \'\<hash\>\',                
                                      \'SHA-512/224\':             
                                      \'\<hash\>\',                
                                      \'SHA-512/256\':             
                                      \'\<hash\>\'}}\"}            

  Report            A report to       {id: 1, event: Report, data: Implemented
                    download          \"{\'url\':                  
                                      \'<http://download.me>\',    
                                      \'name\': \'filename.txt\',  
                                      \'checksums\': {\'SHA-224\': 
                                      \'\<hash\>\', \'SHA-256\':   
                                      \'\<hash\>\', \'SHA-384\':   
                                      \'\<hash\>\', \'SHA-512\':   
                                      \'\<hash\>\',                
                                      \'SHA-512/224\':             
                                      \'\<hash\>\',                
                                      \'SHA-512/256\':             
                                      \'\<hash\>\'}}\"}            

  TestCase          A test case       {id: 1, event: TestCase,     Suggested
                    execution         data: \"{\'id\': \'uuid\',   
                                      \'triggered\': True,         
                                      \'started\': True,           
                                      \'finished\': False}\"}      
  ----------------------------------------------------------------------------------

  : SSE Protocol
