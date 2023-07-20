```html
GOROOT=/usr/local/opt/go/libexec #gosetup
GOPATH=/Users/qun/go #gosetup
/usr/local/opt/go/libexec/bin/go build -o /Users/qun/Library/Caches/JetBrains/GoLand2023.1/tmp/GoLand/___go_build_github_com_houseme_goframe_polaris_demo_http__3_ /Users/qun/Documents/golang/polarismesh/goframe-polaris-demo/http/client.go #gosetup
/Users/qun/Library/Caches/JetBrains/GoLand2023.1/tmp/GoLand/___go_build_github_com_houseme_goframe_polaris_demo_http__3_
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID:  , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID:  , instance.GetId: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID:  , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID:  , instance.GetId: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  0  time:  2023-07-20 20:42:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  1  time:  2023-07-20 20:42:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  2  time:  2023-07-20 20:42:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  3  time:  2023-07-20 20:42:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  4  time:  2023-07-20 20:42:14  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  5  time:  2023-07-20 20:42:15  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  6  time:  2023-07-20 20:42:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  7  time:  2023-07-20 20:42:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  8  time:  2023-07-20 20:42:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  9  time:  2023-07-20 20:42:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  10  time:  2023-07-20 20:42:20  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  11  time:  2023-07-20 20:42:21  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  12  time:  2023-07-20 20:42:22  code:  200
AddEvent start ServiceInstances len: 2 instanceEvent.AddEvent.Instances len: 1
allEndpointStr: 192.168.1.2:63842,192.168.1.2:63853 endpointStr: 192.168.1.2:64140,
newEndpointStr: 192.168.1.2:64140,
AddEvent fmt end allEndpointStr: 192.168.1.2:64140,192.168.1.2:63842,192.168.1.2:63853
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID:  , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
AddEvent end ServiceInstances len: 3
services: [0xc0003a0440 0xc0003a0560 0xc0003a0680]
services[0]: &{Service:0xc000529730 ID:986ed267ca152aff26b767e1c1e2fd4820271f47},endpoint:192.168.1.2:64140,192.168.1.2:63842,192.168.1.2:63853
services[1]: &{Service:0xc0005297a0 ID:a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26},endpoint:192.168.1.2:64140,192.168.1.2:63842,192.168.1.2:63853
services[2]: &{Service:0xc000529810 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:64140,192.168.1.2:63842,192.168.1.2:63853
services[0]: 192.168.1.2:64140,192.168.1.2:63842,192.168.1.2:63853
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  13  time:  2023-07-20 20:42:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  14  time:  2023-07-20 20:42:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  15  time:  2023-07-20 20:42:25  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  16  time:  2023-07-20 20:42:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  17  time:  2023-07-20 20:42:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  18  time:  2023-07-20 20:42:28  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  19  time:  2023-07-20 20:42:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  20  time:  2023-07-20 20:42:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  21  time:  2023-07-20 20:42:31  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  22  time:  2023-07-20 20:42:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  23  time:  2023-07-20 20:42:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  24  time:  2023-07-20 20:42:34  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  25  time:  2023-07-20 20:42:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  26  time:  2023-07-20 20:42:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  27  time:  2023-07-20 20:42:37  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  28  time:  2023-07-20 20:42:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  29  time:  2023-07-20 20:42:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  30  time:  2023-07-20 20:42:40  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  31  time:  2023-07-20 20:42:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  32  time:  2023-07-20 20:42:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  33  time:  2023-07-20 20:42:43  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  34  time:  2023-07-20 20:42:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  35  time:  2023-07-20 20:42:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  36  time:  2023-07-20 20:42:46  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  37  time:  2023-07-20 20:42:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  38  time:  2023-07-20 20:42:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  39  time:  2023-07-20 20:42:49  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  40  time:  2023-07-20 20:42:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  41  time:  2023-07-20 20:42:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  42  time:  2023-07-20 20:42:52  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  43  time:  2023-07-20 20:42:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  44  time:  2023-07-20 20:42:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  45  time:  2023-07-20 20:42:55  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  46  time:  2023-07-20 20:42:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  47  time:  2023-07-20 20:42:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  48  time:  2023-07-20 20:42:58  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  49  time:  2023-07-20 20:42:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  50  time:  2023-07-20 20:43:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  51  time:  2023-07-20 20:43:01  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  52  time:  2023-07-20 20:43:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  53  time:  2023-07-20 20:43:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  54  time:  2023-07-20 20:43:04  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  55  time:  2023-07-20 20:43:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  56  time:  2023-07-20 20:43:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  57  time:  2023-07-20 20:43:07  code:  200
DeleteEvent start ServiceInstances len: 3 instanceEvent.DeleteEvent.Instances len: 1
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID: fba451c4333b44a902cf92a3283c915549ef43fa , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
services: [0xc00007f140 0xc00007f240 0xc00007f340]
services[0]: &{Service:0xc0003228c0 ID:986ed267ca152aff26b767e1c1e2fd4820271f47},endpoint:192.168.1.2:63842,192.168.1.2:63853
services[1]: &{Service:0xc000322930 ID:a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26},endpoint:192.168.1.2:63842,192.168.1.2:63853
services[2]: &{Service:0xc0003229a0 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:63842,192.168.1.2:63853
services[0]: 192.168.1.2:63842,192.168.1.2:63853
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  58  time:  2023-07-20 20:43:08  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  59  time:  2023-07-20 20:43:09  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  60  time:  2023-07-20 20:43:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  61  time:  2023-07-20 20:43:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  62  time:  2023-07-20 20:43:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  63  time:  2023-07-20 20:43:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  64  time:  2023-07-20 20:43:14  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  65  time:  2023-07-20 20:43:15  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  66  time:  2023-07-20 20:43:16  code:  200
UpdateEvent start ServiceInstances len: 3 instanceEvent.UpdateEvent.UpdateList len: 1
UpdateEvent start serviceEndpointStr: 192.168.1.2:63842,192.168.1.2:63853 ,endpointStr: 192.168.1.2:63842, ,healthyEndpointStr: 192.168.1.2:63842,
UpdateEvent SplitAndTrim endpointStr serviceEndpointStr: 192.168.1.2:63842,192.168.1.2:63853 newEndpointStr: 192.168.1.2:63853,
UpdateEvent add healthyEndpointStr serviceEndpointStr: 192.168.1.2:63842,192.168.1.2:63853 newEndpointStr: 192.168.1.2:63853,192.168.1.2:63842,
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID: fba451c4333b44a902cf92a3283c915549ef43fa , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
services: [0xc0003a1720 0xc0003a1820 0xc0003a1920]
services[0]: &{Service:0xc00049b960 ID:986ed267ca152aff26b767e1c1e2fd4820271f47},endpoint:192.168.1.2:63853,192.168.1.2:63842
services[1]: &{Service:0xc00049b9d0 ID:a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26},endpoint:192.168.1.2:63853,192.168.1.2:63842
services[2]: &{Service:0xc00049ba40 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:63853,192.168.1.2:63842
services[0]: 192.168.1.2:63853,192.168.1.2:63842
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  67  time:  2023-07-20 20:43:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  68  time:  2023-07-20 20:43:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  69  time:  2023-07-20 20:43:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  70  time:  2023-07-20 20:43:20  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  71  time:  2023-07-20 20:43:21  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  72  time:  2023-07-20 20:43:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  73  time:  2023-07-20 20:43:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  74  time:  2023-07-20 20:43:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  75  time:  2023-07-20 20:43:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  76  time:  2023-07-20 20:43:26  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  77  time:  2023-07-20 20:43:27  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  78  time:  2023-07-20 20:43:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  79  time:  2023-07-20 20:43:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  80  time:  2023-07-20 20:43:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  81  time:  2023-07-20 20:43:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  82  time:  2023-07-20 20:43:32  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  83  time:  2023-07-20 20:43:33  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  84  time:  2023-07-20 20:43:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  85  time:  2023-07-20 20:43:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  86  time:  2023-07-20 20:43:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  87  time:  2023-07-20 20:43:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  88  time:  2023-07-20 20:43:38  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  89  time:  2023-07-20 20:43:39  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  90  time:  2023-07-20 20:43:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  91  time:  2023-07-20 20:43:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  92  time:  2023-07-20 20:43:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  93  time:  2023-07-20 20:43:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  94  time:  2023-07-20 20:43:44  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  95  time:  2023-07-20 20:43:45  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  96  time:  2023-07-20 20:43:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  97  time:  2023-07-20 20:43:47  code:  200
UpdateEvent start ServiceInstances len: 3 instanceEvent.UpdateEvent.UpdateList len: 1
UpdateEvent start serviceEndpointStr: 192.168.1.2:63853,192.168.1.2:63842 ,endpointStr: 192.168.1.2:63842, ,healthyEndpointStr: 
UpdateEvent SplitAndTrim endpointStr serviceEndpointStr: 192.168.1.2:63853,192.168.1.2:63842 newEndpointStr: 192.168.1.2:63853,
UpdateEvent add healthyEndpointStr serviceEndpointStr: 192.168.1.2:63853,192.168.1.2:63842 newEndpointStr: 192.168.1.2:63853,
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID: fba451c4333b44a902cf92a3283c915549ef43fa , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
services: [0xc000125d00 0xc000125d60 0xc000125dc0]
services[0]: &{Service:0xc00049b5e0 ID:986ed267ca152aff26b767e1c1e2fd4820271f47},endpoint:192.168.1.2:63853
services[1]: &{Service:0xc00049b650 ID:a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26},endpoint:192.168.1.2:63853
services[2]: &{Service:0xc00049b6c0 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:63853
services[0]: 192.168.1.2:63853
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  98  time:  2023-07-20 20:43:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  99  time:  2023-07-20 20:43:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  100  time:  2023-07-20 20:43:50  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  101  time:  2023-07-20 20:43:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  102  time:  2023-07-20 20:43:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  103  time:  2023-07-20 20:43:53  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  104  time:  2023-07-20 20:43:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  105  time:  2023-07-20 20:43:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  106  time:  2023-07-20 20:43:56  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  107  time:  2023-07-20 20:43:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  108  time:  2023-07-20 20:43:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  109  time:  2023-07-20 20:43:59  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  110  time:  2023-07-20 20:44:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  111  time:  2023-07-20 20:44:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  112  time:  2023-07-20 20:44:02  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  113  time:  2023-07-20 20:44:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  114  time:  2023-07-20 20:44:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  115  time:  2023-07-20 20:44:05  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  116  time:  2023-07-20 20:44:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  117  time:  2023-07-20 20:44:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  118  time:  2023-07-20 20:44:08  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  119  time:  2023-07-20 20:44:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  120  time:  2023-07-20 20:44:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  121  time:  2023-07-20 20:44:11  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  122  time:  2023-07-20 20:44:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  123  time:  2023-07-20 20:44:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  124  time:  2023-07-20 20:44:14  code:  200
UpdateEvent start ServiceInstances len: 3 instanceEvent.UpdateEvent.UpdateList len: 1
UpdateEvent start serviceEndpointStr: 192.168.1.2:63853 ,endpointStr: 192.168.1.2:63842, ,healthyEndpointStr: 192.168.1.2:63842,
UpdateEvent SplitAndTrim endpointStr serviceEndpointStr: 192.168.1.2:63853 newEndpointStr: 192.168.1.2:63853,
UpdateEvent add healthyEndpointStr serviceEndpointStr: 192.168.1.2:63853 newEndpointStr: 192.168.1.2:63853,192.168.1.2:63842,
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID: fba451c4333b44a902cf92a3283c915549ef43fa , instance.GetId: 986ed267ca152aff26b767e1c1e2fd4820271f47
AddEvent start ServiceInstances len: 3 instanceEvent.AddEvent.Instances len: 1
allEndpointStr: 192.168.1.2:63853,192.168.1.2:63842 endpointStr: 192.168.1.2:64140,
newEndpointStr: 192.168.1.2:64140,
AddEvent fmt end allEndpointStr: 192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
instanceToServiceInstance 986ed267ca152aff26b767e1c1e2fd4820271f47 , instanceID: 986ed267ca152aff26b767e1c1e2fd4820271f47 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instanceID: a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26 , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID: fba451c4333b44a902cf92a3283c915549ef43fa , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
instanceToServiceInstance fba451c4333b44a902cf92a3283c915549ef43fa , instanceID:  , instance.GetId: fba451c4333b44a902cf92a3283c915549ef43fa
AddEvent end ServiceInstances len: 4
services: [0xc000736dc0 0xc000736ee0 0xc000737000 0xc000737120]
services[0]: &{Service:0xc000322930 ID:986ed267ca152aff26b767e1c1e2fd4820271f47},endpoint:192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
services[1]: &{Service:0xc0003229a0 ID:a2d200de5f985e4938ef142f6bbb5ab1cdcd1e26},endpoint:192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
services[2]: &{Service:0xc000322a10 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
services[3]: &{Service:0xc000322a80 ID:fba451c4333b44a902cf92a3283c915549ef43fa},endpoint:192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
services[0]: 192.168.1.2:64140,192.168.1.2:63853,192.168.1.2:63842
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  125  time:  2023-07-20 20:44:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  126  time:  2023-07-20 20:44:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  127  time:  2023-07-20 20:44:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  128  time:  2023-07-20 20:44:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  129  time:  2023-07-20 20:44:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  130  time:  2023-07-20 20:44:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  131  time:  2023-07-20 20:44:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  132  time:  2023-07-20 20:44:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  133  time:  2023-07-20 20:44:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  134  time:  2023-07-20 20:44:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  135  time:  2023-07-20 20:44:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  136  time:  2023-07-20 20:44:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  137  time:  2023-07-20 20:44:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  138  time:  2023-07-20 20:44:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  139  time:  2023-07-20 20:44:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  140  time:  2023-07-20 20:44:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  141  time:  2023-07-20 20:44:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  142  time:  2023-07-20 20:44:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  143  time:  2023-07-20 20:44:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  144  time:  2023-07-20 20:44:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  145  time:  2023-07-20 20:44:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  146  time:  2023-07-20 20:44:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  147  time:  2023-07-20 20:44:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  148  time:  2023-07-20 20:44:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  149  time:  2023-07-20 20:44:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  150  time:  2023-07-20 20:44:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  151  time:  2023-07-20 20:44:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  152  time:  2023-07-20 20:44:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  153  time:  2023-07-20 20:44:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  154  time:  2023-07-20 20:44:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  155  time:  2023-07-20 20:44:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  156  time:  2023-07-20 20:44:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  157  time:  2023-07-20 20:44:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  158  time:  2023-07-20 20:44:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  159  time:  2023-07-20 20:44:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  160  time:  2023-07-20 20:44:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  161  time:  2023-07-20 20:44:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  162  time:  2023-07-20 20:44:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  163  time:  2023-07-20 20:44:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  164  time:  2023-07-20 20:44:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  165  time:  2023-07-20 20:44:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  166  time:  2023-07-20 20:44:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  167  time:  2023-07-20 20:44:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  168  time:  2023-07-20 20:44:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  169  time:  2023-07-20 20:44:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  170  time:  2023-07-20 20:45:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  171  time:  2023-07-20 20:45:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  172  time:  2023-07-20 20:45:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  173  time:  2023-07-20 20:45:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  174  time:  2023-07-20 20:45:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  175  time:  2023-07-20 20:45:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  176  time:  2023-07-20 20:45:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  177  time:  2023-07-20 20:45:07  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  178  time:  2023-07-20 20:45:08  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  179  time:  2023-07-20 20:45:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  180  time:  2023-07-20 20:45:10  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  181  time:  2023-07-20 20:45:11  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  182  time:  2023-07-20 20:45:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  183  time:  2023-07-20 20:45:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  184  time:  2023-07-20 20:45:14  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  185  time:  2023-07-20 20:45:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  186  time:  2023-07-20 20:45:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  187  time:  2023-07-20 20:45:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  188  time:  2023-07-20 20:45:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  189  time:  2023-07-20 20:45:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  190  time:  2023-07-20 20:45:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  191  time:  2023-07-20 20:45:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  192  time:  2023-07-20 20:45:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  193  time:  2023-07-20 20:45:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  194  time:  2023-07-20 20:45:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  195  time:  2023-07-20 20:45:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  196  time:  2023-07-20 20:45:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  197  time:  2023-07-20 20:45:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  198  time:  2023-07-20 20:45:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  199  time:  2023-07-20 20:45:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  200  time:  2023-07-20 20:45:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  201  time:  2023-07-20 20:45:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  202  time:  2023-07-20 20:45:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  203  time:  2023-07-20 20:45:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  204  time:  2023-07-20 20:45:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  205  time:  2023-07-20 20:45:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  206  time:  2023-07-20 20:45:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  207  time:  2023-07-20 20:45:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  208  time:  2023-07-20 20:45:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  209  time:  2023-07-20 20:45:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  210  time:  2023-07-20 20:45:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  211  time:  2023-07-20 20:45:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  212  time:  2023-07-20 20:45:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  213  time:  2023-07-20 20:45:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  214  time:  2023-07-20 20:45:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  215  time:  2023-07-20 20:45:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  216  time:  2023-07-20 20:45:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  217  time:  2023-07-20 20:45:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  218  time:  2023-07-20 20:45:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  219  time:  2023-07-20 20:45:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  220  time:  2023-07-20 20:45:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  221  time:  2023-07-20 20:45:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  222  time:  2023-07-20 20:45:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  223  time:  2023-07-20 20:45:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  224  time:  2023-07-20 20:45:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  225  time:  2023-07-20 20:45:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  226  time:  2023-07-20 20:45:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  227  time:  2023-07-20 20:45:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  228  time:  2023-07-20 20:45:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  229  time:  2023-07-20 20:45:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  230  time:  2023-07-20 20:46:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  231  time:  2023-07-20 20:46:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  232  time:  2023-07-20 20:46:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  233  time:  2023-07-20 20:46:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  234  time:  2023-07-20 20:46:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  235  time:  2023-07-20 20:46:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  236  time:  2023-07-20 20:46:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  237  time:  2023-07-20 20:46:07  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  238  time:  2023-07-20 20:46:08  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  239  time:  2023-07-20 20:46:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  240  time:  2023-07-20 20:46:10  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  241  time:  2023-07-20 20:46:11  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  242  time:  2023-07-20 20:46:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  243  time:  2023-07-20 20:46:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  244  time:  2023-07-20 20:46:14  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  245  time:  2023-07-20 20:46:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  246  time:  2023-07-20 20:46:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  247  time:  2023-07-20 20:46:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  248  time:  2023-07-20 20:46:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  249  time:  2023-07-20 20:46:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  250  time:  2023-07-20 20:46:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  251  time:  2023-07-20 20:46:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  252  time:  2023-07-20 20:46:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  253  time:  2023-07-20 20:46:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  254  time:  2023-07-20 20:46:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  255  time:  2023-07-20 20:46:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  256  time:  2023-07-20 20:46:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  257  time:  2023-07-20 20:46:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  258  time:  2023-07-20 20:46:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  259  time:  2023-07-20 20:46:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  260  time:  2023-07-20 20:46:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  261  time:  2023-07-20 20:46:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  262  time:  2023-07-20 20:46:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  263  time:  2023-07-20 20:46:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  264  time:  2023-07-20 20:46:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  265  time:  2023-07-20 20:46:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  266  time:  2023-07-20 20:46:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  267  time:  2023-07-20 20:46:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  268  time:  2023-07-20 20:46:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  269  time:  2023-07-20 20:46:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  270  time:  2023-07-20 20:46:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  271  time:  2023-07-20 20:46:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  272  time:  2023-07-20 20:46:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  273  time:  2023-07-20 20:46:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  274  time:  2023-07-20 20:46:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  275  time:  2023-07-20 20:46:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  276  time:  2023-07-20 20:46:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  277  time:  2023-07-20 20:46:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  278  time:  2023-07-20 20:46:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  279  time:  2023-07-20 20:46:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  280  time:  2023-07-20 20:46:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  281  time:  2023-07-20 20:46:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  282  time:  2023-07-20 20:46:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  283  time:  2023-07-20 20:46:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  284  time:  2023-07-20 20:46:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  285  time:  2023-07-20 20:46:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  286  time:  2023-07-20 20:46:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  287  time:  2023-07-20 20:46:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  288  time:  2023-07-20 20:46:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  289  time:  2023-07-20 20:46:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  290  time:  2023-07-20 20:47:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  291  time:  2023-07-20 20:47:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  292  time:  2023-07-20 20:47:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  293  time:  2023-07-20 20:47:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  294  time:  2023-07-20 20:47:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  295  time:  2023-07-20 20:47:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  296  time:  2023-07-20 20:47:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  297  time:  2023-07-20 20:47:07  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  298  time:  2023-07-20 20:47:08  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  299  time:  2023-07-20 20:47:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  300  time:  2023-07-20 20:47:10  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  301  time:  2023-07-20 20:47:11  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  302  time:  2023-07-20 20:47:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  303  time:  2023-07-20 20:47:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  304  time:  2023-07-20 20:47:14  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  305  time:  2023-07-20 20:47:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  306  time:  2023-07-20 20:47:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  307  time:  2023-07-20 20:47:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  308  time:  2023-07-20 20:47:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  309  time:  2023-07-20 20:47:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  310  time:  2023-07-20 20:47:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  311  time:  2023-07-20 20:47:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  312  time:  2023-07-20 20:47:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  313  time:  2023-07-20 20:47:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  314  time:  2023-07-20 20:47:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  315  time:  2023-07-20 20:47:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  316  time:  2023-07-20 20:47:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  317  time:  2023-07-20 20:47:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  318  time:  2023-07-20 20:47:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  319  time:  2023-07-20 20:47:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  320  time:  2023-07-20 20:47:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  321  time:  2023-07-20 20:47:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  322  time:  2023-07-20 20:47:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  323  time:  2023-07-20 20:47:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  324  time:  2023-07-20 20:47:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  325  time:  2023-07-20 20:47:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  326  time:  2023-07-20 20:47:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  327  time:  2023-07-20 20:47:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  328  time:  2023-07-20 20:47:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  329  time:  2023-07-20 20:47:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  330  time:  2023-07-20 20:47:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  331  time:  2023-07-20 20:47:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  332  time:  2023-07-20 20:47:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  333  time:  2023-07-20 20:47:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  334  time:  2023-07-20 20:47:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  335  time:  2023-07-20 20:47:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  336  time:  2023-07-20 20:47:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  337  time:  2023-07-20 20:47:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  338  time:  2023-07-20 20:47:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  339  time:  2023-07-20 20:47:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  340  time:  2023-07-20 20:47:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  341  time:  2023-07-20 20:47:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  342  time:  2023-07-20 20:47:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  343  time:  2023-07-20 20:47:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  344  time:  2023-07-20 20:47:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  345  time:  2023-07-20 20:47:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  346  time:  2023-07-20 20:47:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  347  time:  2023-07-20 20:47:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  348  time:  2023-07-20 20:47:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  349  time:  2023-07-20 20:47:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  350  time:  2023-07-20 20:48:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  351  time:  2023-07-20 20:48:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  352  time:  2023-07-20 20:48:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  353  time:  2023-07-20 20:48:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  354  time:  2023-07-20 20:48:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  355  time:  2023-07-20 20:48:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  356  time:  2023-07-20 20:48:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  357  time:  2023-07-20 20:48:07  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  358  time:  2023-07-20 20:48:08  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  359  time:  2023-07-20 20:48:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  360  time:  2023-07-20 20:48:10  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  361  time:  2023-07-20 20:48:11  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  362  time:  2023-07-20 20:48:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  363  time:  2023-07-20 20:48:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  364  time:  2023-07-20 20:48:14  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  365  time:  2023-07-20 20:48:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  366  time:  2023-07-20 20:48:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  367  time:  2023-07-20 20:48:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  368  time:  2023-07-20 20:48:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  369  time:  2023-07-20 20:48:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  370  time:  2023-07-20 20:48:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  371  time:  2023-07-20 20:48:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  372  time:  2023-07-20 20:48:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  373  time:  2023-07-20 20:48:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  374  time:  2023-07-20 20:48:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  375  time:  2023-07-20 20:48:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  376  time:  2023-07-20 20:48:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  377  time:  2023-07-20 20:48:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  378  time:  2023-07-20 20:48:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  379  time:  2023-07-20 20:48:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  380  time:  2023-07-20 20:48:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  381  time:  2023-07-20 20:48:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  382  time:  2023-07-20 20:48:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  383  time:  2023-07-20 20:48:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  384  time:  2023-07-20 20:48:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  385  time:  2023-07-20 20:48:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  386  time:  2023-07-20 20:48:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  387  time:  2023-07-20 20:48:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  388  time:  2023-07-20 20:48:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  389  time:  2023-07-20 20:48:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  390  time:  2023-07-20 20:48:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  391  time:  2023-07-20 20:48:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  392  time:  2023-07-20 20:48:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  393  time:  2023-07-20 20:48:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  394  time:  2023-07-20 20:48:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  395  time:  2023-07-20 20:48:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  396  time:  2023-07-20 20:48:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  397  time:  2023-07-20 20:48:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  398  time:  2023-07-20 20:48:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  399  time:  2023-07-20 20:48:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  400  time:  2023-07-20 20:48:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  401  time:  2023-07-20 20:48:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  402  time:  2023-07-20 20:48:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  403  time:  2023-07-20 20:48:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  404  time:  2023-07-20 20:48:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  405  time:  2023-07-20 20:48:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  406  time:  2023-07-20 20:48:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  407  time:  2023-07-20 20:48:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  408  time:  2023-07-20 20:48:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  409  time:  2023-07-20 20:48:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  410  time:  2023-07-20 20:49:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  411  time:  2023-07-20 20:49:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  412  time:  2023-07-20 20:49:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  413  time:  2023-07-20 20:49:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  414  time:  2023-07-20 20:49:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  415  time:  2023-07-20 20:49:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  416  time:  2023-07-20 20:49:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  417  time:  2023-07-20 20:49:07  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  418  time:  2023-07-20 20:49:08  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  419  time:  2023-07-20 20:49:09  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  420  time:  2023-07-20 20:49:10  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  421  time:  2023-07-20 20:49:11  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  422  time:  2023-07-20 20:49:12  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  423  time:  2023-07-20 20:49:13  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  424  time:  2023-07-20 20:49:14  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  425  time:  2023-07-20 20:49:15  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  426  time:  2023-07-20 20:49:16  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  427  time:  2023-07-20 20:49:17  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  428  time:  2023-07-20 20:49:18  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  429  time:  2023-07-20 20:49:19  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  430  time:  2023-07-20 20:49:20  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  431  time:  2023-07-20 20:49:21  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  432  time:  2023-07-20 20:49:22  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  433  time:  2023-07-20 20:49:23  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  434  time:  2023-07-20 20:49:24  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  435  time:  2023-07-20 20:49:25  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  436  time:  2023-07-20 20:49:26  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  437  time:  2023-07-20 20:49:27  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  438  time:  2023-07-20 20:49:28  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  439  time:  2023-07-20 20:49:29  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  440  time:  2023-07-20 20:49:30  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  441  time:  2023-07-20 20:49:31  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  442  time:  2023-07-20 20:49:32  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  443  time:  2023-07-20 20:49:33  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  444  time:  2023-07-20 20:49:34  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  445  time:  2023-07-20 20:49:35  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  446  time:  2023-07-20 20:49:36  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  447  time:  2023-07-20 20:49:37  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  448  time:  2023-07-20 20:49:38  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  449  time:  2023-07-20 20:49:39  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  450  time:  2023-07-20 20:49:40  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  451  time:  2023-07-20 20:49:41  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  452  time:  2023-07-20 20:49:42  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  453  time:  2023-07-20 20:49:43  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  454  time:  2023-07-20 20:49:44  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  455  time:  2023-07-20 20:49:45  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  456  time:  2023-07-20 20:49:46  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  457  time:  2023-07-20 20:49:47  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  458  time:  2023-07-20 20:49:48  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  459  time:  2023-07-20 20:49:49  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  460  time:  2023-07-20 20:49:50  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  461  time:  2023-07-20 20:49:51  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  462  time:  2023-07-20 20:49:52  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  463  time:  2023-07-20 20:49:53  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  464  time:  2023-07-20 20:49:54  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  465  time:  2023-07-20 20:49:55  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  466  time:  2023-07-20 20:49:56  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  467  time:  2023-07-20 20:49:57  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  468  time:  2023-07-20 20:49:58  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  469  time:  2023-07-20 20:49:59  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  470  time:  2023-07-20 20:50:00  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  471  time:  2023-07-20 20:50:01  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  472  time:  2023-07-20 20:50:02  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  473  time:  2023-07-20 20:50:03  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  474  time:  2023-07-20 20:50:04  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  475  time:  2023-07-20 20:50:05  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  476  time:  2023-07-20 20:50:06  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  477  time:  2023-07-20 20:50:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  478  time:  2023-07-20 20:50:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  479  time:  2023-07-20 20:50:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  480  time:  2023-07-20 20:50:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  481  time:  2023-07-20 20:50:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  482  time:  2023-07-20 20:50:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  483  time:  2023-07-20 20:50:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  484  time:  2023-07-20 20:50:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  485  time:  2023-07-20 20:50:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  486  time:  2023-07-20 20:50:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  487  time:  2023-07-20 20:50:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  488  time:  2023-07-20 20:50:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  489  time:  2023-07-20 20:50:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  490  time:  2023-07-20 20:50:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  491  time:  2023-07-20 20:50:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  492  time:  2023-07-20 20:50:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  493  time:  2023-07-20 20:50:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  494  time:  2023-07-20 20:50:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  495  time:  2023-07-20 20:50:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  496  time:  2023-07-20 20:50:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  497  time:  2023-07-20 20:50:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  498  time:  2023-07-20 20:50:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  499  time:  2023-07-20 20:50:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  500  time:  2023-07-20 20:50:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  501  time:  2023-07-20 20:50:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  502  time:  2023-07-20 20:50:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  503  time:  2023-07-20 20:50:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  504  time:  2023-07-20 20:50:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  505  time:  2023-07-20 20:50:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  506  time:  2023-07-20 20:50:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  507  time:  2023-07-20 20:50:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  508  time:  2023-07-20 20:50:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  509  time:  2023-07-20 20:50:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  510  time:  2023-07-20 20:50:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  511  time:  2023-07-20 20:50:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  512  time:  2023-07-20 20:50:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  513  time:  2023-07-20 20:50:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  514  time:  2023-07-20 20:50:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  515  time:  2023-07-20 20:50:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  516  time:  2023-07-20 20:50:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  517  time:  2023-07-20 20:50:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  518  time:  2023-07-20 20:50:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  519  time:  2023-07-20 20:50:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  520  time:  2023-07-20 20:50:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  521  time:  2023-07-20 20:50:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  522  time:  2023-07-20 20:50:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  523  time:  2023-07-20 20:50:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  524  time:  2023-07-20 20:50:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  525  time:  2023-07-20 20:50:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  526  time:  2023-07-20 20:50:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  527  time:  2023-07-20 20:50:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  528  time:  2023-07-20 20:50:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  529  time:  2023-07-20 20:51:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  530  time:  2023-07-20 20:51:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  531  time:  2023-07-20 20:51:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  532  time:  2023-07-20 20:51:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  533  time:  2023-07-20 20:51:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  534  time:  2023-07-20 20:51:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  535  time:  2023-07-20 20:51:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  536  time:  2023-07-20 20:51:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  537  time:  2023-07-20 20:51:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  538  time:  2023-07-20 20:51:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  539  time:  2023-07-20 20:51:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  540  time:  2023-07-20 20:51:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  541  time:  2023-07-20 20:51:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  542  time:  2023-07-20 20:51:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  543  time:  2023-07-20 20:51:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  544  time:  2023-07-20 20:51:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  545  time:  2023-07-20 20:51:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  546  time:  2023-07-20 20:51:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  547  time:  2023-07-20 20:51:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  548  time:  2023-07-20 20:51:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  549  time:  2023-07-20 20:51:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  550  time:  2023-07-20 20:51:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  551  time:  2023-07-20 20:51:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  552  time:  2023-07-20 20:51:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  553  time:  2023-07-20 20:51:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  554  time:  2023-07-20 20:51:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  555  time:  2023-07-20 20:51:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  556  time:  2023-07-20 20:51:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  557  time:  2023-07-20 20:51:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  558  time:  2023-07-20 20:51:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  559  time:  2023-07-20 20:51:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  560  time:  2023-07-20 20:51:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  561  time:  2023-07-20 20:51:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  562  time:  2023-07-20 20:51:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  563  time:  2023-07-20 20:51:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  564  time:  2023-07-20 20:51:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  565  time:  2023-07-20 20:51:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  566  time:  2023-07-20 20:51:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  567  time:  2023-07-20 20:51:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  568  time:  2023-07-20 20:51:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  569  time:  2023-07-20 20:51:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  570  time:  2023-07-20 20:51:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  571  time:  2023-07-20 20:51:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  572  time:  2023-07-20 20:51:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  573  time:  2023-07-20 20:51:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  574  time:  2023-07-20 20:51:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  575  time:  2023-07-20 20:51:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  576  time:  2023-07-20 20:51:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  577  time:  2023-07-20 20:51:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  578  time:  2023-07-20 20:51:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  579  time:  2023-07-20 20:51:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  580  time:  2023-07-20 20:51:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  581  time:  2023-07-20 20:51:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  582  time:  2023-07-20 20:51:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  583  time:  2023-07-20 20:51:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  584  time:  2023-07-20 20:51:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  585  time:  2023-07-20 20:51:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  586  time:  2023-07-20 20:51:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  587  time:  2023-07-20 20:51:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  588  time:  2023-07-20 20:51:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  589  time:  2023-07-20 20:52:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  590  time:  2023-07-20 20:52:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  591  time:  2023-07-20 20:52:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  592  time:  2023-07-20 20:52:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  593  time:  2023-07-20 20:52:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  594  time:  2023-07-20 20:52:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  595  time:  2023-07-20 20:52:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  596  time:  2023-07-20 20:52:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  597  time:  2023-07-20 20:52:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  598  time:  2023-07-20 20:52:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  599  time:  2023-07-20 20:52:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  600  time:  2023-07-20 20:52:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  601  time:  2023-07-20 20:52:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  602  time:  2023-07-20 20:52:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  603  time:  2023-07-20 20:52:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  604  time:  2023-07-20 20:52:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  605  time:  2023-07-20 20:52:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  606  time:  2023-07-20 20:52:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  607  time:  2023-07-20 20:52:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  608  time:  2023-07-20 20:52:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  609  time:  2023-07-20 20:52:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  610  time:  2023-07-20 20:52:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  611  time:  2023-07-20 20:52:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  612  time:  2023-07-20 20:52:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  613  time:  2023-07-20 20:52:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  614  time:  2023-07-20 20:52:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  615  time:  2023-07-20 20:52:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  616  time:  2023-07-20 20:52:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  617  time:  2023-07-20 20:52:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  618  time:  2023-07-20 20:52:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  619  time:  2023-07-20 20:52:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  620  time:  2023-07-20 20:52:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  621  time:  2023-07-20 20:52:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  622  time:  2023-07-20 20:52:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  623  time:  2023-07-20 20:52:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  624  time:  2023-07-20 20:52:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  625  time:  2023-07-20 20:52:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  626  time:  2023-07-20 20:52:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  627  time:  2023-07-20 20:52:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  628  time:  2023-07-20 20:52:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  629  time:  2023-07-20 20:52:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  630  time:  2023-07-20 20:52:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  631  time:  2023-07-20 20:52:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  632  time:  2023-07-20 20:52:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  633  time:  2023-07-20 20:52:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  634  time:  2023-07-20 20:52:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  635  time:  2023-07-20 20:52:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  636  time:  2023-07-20 20:52:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  637  time:  2023-07-20 20:52:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  638  time:  2023-07-20 20:52:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  639  time:  2023-07-20 20:52:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  640  time:  2023-07-20 20:52:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  641  time:  2023-07-20 20:52:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  642  time:  2023-07-20 20:52:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  643  time:  2023-07-20 20:52:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  644  time:  2023-07-20 20:52:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  645  time:  2023-07-20 20:52:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  646  time:  2023-07-20 20:52:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  647  time:  2023-07-20 20:52:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  648  time:  2023-07-20 20:52:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  649  time:  2023-07-20 20:53:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  650  time:  2023-07-20 20:53:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  651  time:  2023-07-20 20:53:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  652  time:  2023-07-20 20:53:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  653  time:  2023-07-20 20:53:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  654  time:  2023-07-20 20:53:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  655  time:  2023-07-20 20:53:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  656  time:  2023-07-20 20:53:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  657  time:  2023-07-20 20:53:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  658  time:  2023-07-20 20:53:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  659  time:  2023-07-20 20:53:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  660  time:  2023-07-20 20:53:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  661  time:  2023-07-20 20:53:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  662  time:  2023-07-20 20:53:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  663  time:  2023-07-20 20:53:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  664  time:  2023-07-20 20:53:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  665  time:  2023-07-20 20:53:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  666  time:  2023-07-20 20:53:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  667  time:  2023-07-20 20:53:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  668  time:  2023-07-20 20:53:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  669  time:  2023-07-20 20:53:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  670  time:  2023-07-20 20:53:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  671  time:  2023-07-20 20:53:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  672  time:  2023-07-20 20:53:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  673  time:  2023-07-20 20:53:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  674  time:  2023-07-20 20:53:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  675  time:  2023-07-20 20:53:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  676  time:  2023-07-20 20:53:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  677  time:  2023-07-20 20:53:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  678  time:  2023-07-20 20:53:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  679  time:  2023-07-20 20:53:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  680  time:  2023-07-20 20:53:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  681  time:  2023-07-20 20:53:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  682  time:  2023-07-20 20:53:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  683  time:  2023-07-20 20:53:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  684  time:  2023-07-20 20:53:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  685  time:  2023-07-20 20:53:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  686  time:  2023-07-20 20:53:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  687  time:  2023-07-20 20:53:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  688  time:  2023-07-20 20:53:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  689  time:  2023-07-20 20:53:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  690  time:  2023-07-20 20:53:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  691  time:  2023-07-20 20:53:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  692  time:  2023-07-20 20:53:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  693  time:  2023-07-20 20:53:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  694  time:  2023-07-20 20:53:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  695  time:  2023-07-20 20:53:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  696  time:  2023-07-20 20:53:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  697  time:  2023-07-20 20:53:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  698  time:  2023-07-20 20:53:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  699  time:  2023-07-20 20:53:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  700  time:  2023-07-20 20:53:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  701  time:  2023-07-20 20:53:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  702  time:  2023-07-20 20:53:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  703  time:  2023-07-20 20:53:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  704  time:  2023-07-20 20:53:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  705  time:  2023-07-20 20:53:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  706  time:  2023-07-20 20:53:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  707  time:  2023-07-20 20:53:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  708  time:  2023-07-20 20:53:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  709  time:  2023-07-20 20:54:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  710  time:  2023-07-20 20:54:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  711  time:  2023-07-20 20:54:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  712  time:  2023-07-20 20:54:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  713  time:  2023-07-20 20:54:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  714  time:  2023-07-20 20:54:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  715  time:  2023-07-20 20:54:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  716  time:  2023-07-20 20:54:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  717  time:  2023-07-20 20:54:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  718  time:  2023-07-20 20:54:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  719  time:  2023-07-20 20:54:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  720  time:  2023-07-20 20:54:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  721  time:  2023-07-20 20:54:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  722  time:  2023-07-20 20:54:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  723  time:  2023-07-20 20:54:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  724  time:  2023-07-20 20:54:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  725  time:  2023-07-20 20:54:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  726  time:  2023-07-20 20:54:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  727  time:  2023-07-20 20:54:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  728  time:  2023-07-20 20:54:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  729  time:  2023-07-20 20:54:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  730  time:  2023-07-20 20:54:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  731  time:  2023-07-20 20:54:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  732  time:  2023-07-20 20:54:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  733  time:  2023-07-20 20:54:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  734  time:  2023-07-20 20:54:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  735  time:  2023-07-20 20:54:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  736  time:  2023-07-20 20:54:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  737  time:  2023-07-20 20:54:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  738  time:  2023-07-20 20:54:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  739  time:  2023-07-20 20:54:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  740  time:  2023-07-20 20:54:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  741  time:  2023-07-20 20:54:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  742  time:  2023-07-20 20:54:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  743  time:  2023-07-20 20:54:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  744  time:  2023-07-20 20:54:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  745  time:  2023-07-20 20:54:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  746  time:  2023-07-20 20:54:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  747  time:  2023-07-20 20:54:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  748  time:  2023-07-20 20:54:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  749  time:  2023-07-20 20:54:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  750  time:  2023-07-20 20:54:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  751  time:  2023-07-20 20:54:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  752  time:  2023-07-20 20:54:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  753  time:  2023-07-20 20:54:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  754  time:  2023-07-20 20:54:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  755  time:  2023-07-20 20:54:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  756  time:  2023-07-20 20:54:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  757  time:  2023-07-20 20:54:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  758  time:  2023-07-20 20:54:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  759  time:  2023-07-20 20:54:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  760  time:  2023-07-20 20:54:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  761  time:  2023-07-20 20:54:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  762  time:  2023-07-20 20:54:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  763  time:  2023-07-20 20:54:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  764  time:  2023-07-20 20:54:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  765  time:  2023-07-20 20:54:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  766  time:  2023-07-20 20:54:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  767  time:  2023-07-20 20:54:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  768  time:  2023-07-20 20:54:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  769  time:  2023-07-20 20:55:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  770  time:  2023-07-20 20:55:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  771  time:  2023-07-20 20:55:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  772  time:  2023-07-20 20:55:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  773  time:  2023-07-20 20:55:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  774  time:  2023-07-20 20:55:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  775  time:  2023-07-20 20:55:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  776  time:  2023-07-20 20:55:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  777  time:  2023-07-20 20:55:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  778  time:  2023-07-20 20:55:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  779  time:  2023-07-20 20:55:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  780  time:  2023-07-20 20:55:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  781  time:  2023-07-20 20:55:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  782  time:  2023-07-20 20:55:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  783  time:  2023-07-20 20:55:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  784  time:  2023-07-20 20:55:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  785  time:  2023-07-20 20:55:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  786  time:  2023-07-20 20:55:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  787  time:  2023-07-20 20:55:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  788  time:  2023-07-20 20:55:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  789  time:  2023-07-20 20:55:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  790  time:  2023-07-20 20:55:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  791  time:  2023-07-20 20:55:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  792  time:  2023-07-20 20:55:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  793  time:  2023-07-20 20:55:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  794  time:  2023-07-20 20:55:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  795  time:  2023-07-20 20:55:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  796  time:  2023-07-20 20:55:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  797  time:  2023-07-20 20:55:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  798  time:  2023-07-20 20:55:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  799  time:  2023-07-20 20:55:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  800  time:  2023-07-20 20:55:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  801  time:  2023-07-20 20:55:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  802  time:  2023-07-20 20:55:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  803  time:  2023-07-20 20:55:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  804  time:  2023-07-20 20:55:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  805  time:  2023-07-20 20:55:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  806  time:  2023-07-20 20:55:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  807  time:  2023-07-20 20:55:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  808  time:  2023-07-20 20:55:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  809  time:  2023-07-20 20:55:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  810  time:  2023-07-20 20:55:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  811  time:  2023-07-20 20:55:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  812  time:  2023-07-20 20:55:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  813  time:  2023-07-20 20:55:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  814  time:  2023-07-20 20:55:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  815  time:  2023-07-20 20:55:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  816  time:  2023-07-20 20:55:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  817  time:  2023-07-20 20:55:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  818  time:  2023-07-20 20:55:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  819  time:  2023-07-20 20:55:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  820  time:  2023-07-20 20:55:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  821  time:  2023-07-20 20:55:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  822  time:  2023-07-20 20:55:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  823  time:  2023-07-20 20:55:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  824  time:  2023-07-20 20:55:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  825  time:  2023-07-20 20:55:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  826  time:  2023-07-20 20:55:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  827  time:  2023-07-20 20:55:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  828  time:  2023-07-20 20:55:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  829  time:  2023-07-20 20:56:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  830  time:  2023-07-20 20:56:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  831  time:  2023-07-20 20:56:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  832  time:  2023-07-20 20:56:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  833  time:  2023-07-20 20:56:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  834  time:  2023-07-20 20:56:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  835  time:  2023-07-20 20:56:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  836  time:  2023-07-20 20:56:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  837  time:  2023-07-20 20:56:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  838  time:  2023-07-20 20:56:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  839  time:  2023-07-20 20:56:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  840  time:  2023-07-20 20:56:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  841  time:  2023-07-20 20:56:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  842  time:  2023-07-20 20:56:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  843  time:  2023-07-20 20:56:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  844  time:  2023-07-20 20:56:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  845  time:  2023-07-20 20:56:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  846  time:  2023-07-20 20:56:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  847  time:  2023-07-20 20:56:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  848  time:  2023-07-20 20:56:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  849  time:  2023-07-20 20:56:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  850  time:  2023-07-20 20:56:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  851  time:  2023-07-20 20:56:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  852  time:  2023-07-20 20:56:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  853  time:  2023-07-20 20:56:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  854  time:  2023-07-20 20:56:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  855  time:  2023-07-20 20:56:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  856  time:  2023-07-20 20:56:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  857  time:  2023-07-20 20:56:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  858  time:  2023-07-20 20:56:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  859  time:  2023-07-20 20:56:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  860  time:  2023-07-20 20:56:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  861  time:  2023-07-20 20:56:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  862  time:  2023-07-20 20:56:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  863  time:  2023-07-20 20:56:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  864  time:  2023-07-20 20:56:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  865  time:  2023-07-20 20:56:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  866  time:  2023-07-20 20:56:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  867  time:  2023-07-20 20:56:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  868  time:  2023-07-20 20:56:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  869  time:  2023-07-20 20:56:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  870  time:  2023-07-20 20:56:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  871  time:  2023-07-20 20:56:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  872  time:  2023-07-20 20:56:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  873  time:  2023-07-20 20:56:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  874  time:  2023-07-20 20:56:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  875  time:  2023-07-20 20:56:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  876  time:  2023-07-20 20:56:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  877  time:  2023-07-20 20:56:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  878  time:  2023-07-20 20:56:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  879  time:  2023-07-20 20:56:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  880  time:  2023-07-20 20:56:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  881  time:  2023-07-20 20:56:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  882  time:  2023-07-20 20:56:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  883  time:  2023-07-20 20:56:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  884  time:  2023-07-20 20:56:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  885  time:  2023-07-20 20:56:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  886  time:  2023-07-20 20:56:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  887  time:  2023-07-20 20:56:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  888  time:  2023-07-20 20:56:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  889  time:  2023-07-20 20:57:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  890  time:  2023-07-20 20:57:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  891  time:  2023-07-20 20:57:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  892  time:  2023-07-20 20:57:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  893  time:  2023-07-20 20:57:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  894  time:  2023-07-20 20:57:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  895  time:  2023-07-20 20:57:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  896  time:  2023-07-20 20:57:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  897  time:  2023-07-20 20:57:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  898  time:  2023-07-20 20:57:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  899  time:  2023-07-20 20:57:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  900  time:  2023-07-20 20:57:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  901  time:  2023-07-20 20:57:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  902  time:  2023-07-20 20:57:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  903  time:  2023-07-20 20:57:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  904  time:  2023-07-20 20:57:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  905  time:  2023-07-20 20:57:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  906  time:  2023-07-20 20:57:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  907  time:  2023-07-20 20:57:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  908  time:  2023-07-20 20:57:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  909  time:  2023-07-20 20:57:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  910  time:  2023-07-20 20:57:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  911  time:  2023-07-20 20:57:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  912  time:  2023-07-20 20:57:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  913  time:  2023-07-20 20:57:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  914  time:  2023-07-20 20:57:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  915  time:  2023-07-20 20:57:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  916  time:  2023-07-20 20:57:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  917  time:  2023-07-20 20:57:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  918  time:  2023-07-20 20:57:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  919  time:  2023-07-20 20:57:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  920  time:  2023-07-20 20:57:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  921  time:  2023-07-20 20:57:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  922  time:  2023-07-20 20:57:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  923  time:  2023-07-20 20:57:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  924  time:  2023-07-20 20:57:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  925  time:  2023-07-20 20:57:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  926  time:  2023-07-20 20:57:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  927  time:  2023-07-20 20:57:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  928  time:  2023-07-20 20:57:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  929  time:  2023-07-20 20:57:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  930  time:  2023-07-20 20:57:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  931  time:  2023-07-20 20:57:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  932  time:  2023-07-20 20:57:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  933  time:  2023-07-20 20:57:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  934  time:  2023-07-20 20:57:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  935  time:  2023-07-20 20:57:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  936  time:  2023-07-20 20:57:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  937  time:  2023-07-20 20:57:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  938  time:  2023-07-20 20:57:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  939  time:  2023-07-20 20:57:50  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  940  time:  2023-07-20 20:57:51  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  941  time:  2023-07-20 20:57:52  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  942  time:  2023-07-20 20:57:53  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  943  time:  2023-07-20 20:57:54  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  944  time:  2023-07-20 20:57:55  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  945  time:  2023-07-20 20:57:56  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  946  time:  2023-07-20 20:57:57  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  947  time:  2023-07-20 20:57:58  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  948  time:  2023-07-20 20:57:59  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  949  time:  2023-07-20 20:58:00  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  950  time:  2023-07-20 20:58:01  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  951  time:  2023-07-20 20:58:02  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  952  time:  2023-07-20 20:58:03  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  953  time:  2023-07-20 20:58:04  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  954  time:  2023-07-20 20:58:05  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  955  time:  2023-07-20 20:58:06  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  956  time:  2023-07-20 20:58:07  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  957  time:  2023-07-20 20:58:08  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  958  time:  2023-07-20 20:58:09  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  959  time:  2023-07-20 20:58:10  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  960  time:  2023-07-20 20:58:11  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  961  time:  2023-07-20 20:58:12  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  962  time:  2023-07-20 20:58:13  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  963  time:  2023-07-20 20:58:14  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  964  time:  2023-07-20 20:58:15  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  965  time:  2023-07-20 20:58:16  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  966  time:  2023-07-20 20:58:17  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  967  time:  2023-07-20 20:58:18  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  968  time:  2023-07-20 20:58:19  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  969  time:  2023-07-20 20:58:20  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  970  time:  2023-07-20 20:58:21  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  971  time:  2023-07-20 20:58:22  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  972  time:  2023-07-20 20:58:23  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  973  time:  2023-07-20 20:58:24  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  974  time:  2023-07-20 20:58:25  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  975  time:  2023-07-20 20:58:26  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  976  time:  2023-07-20 20:58:27  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  977  time:  2023-07-20 20:58:28  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  978  time:  2023-07-20 20:58:29  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  979  time:  2023-07-20 20:58:30  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  980  time:  2023-07-20 20:58:31  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  981  time:  2023-07-20 20:58:32  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  982  time:  2023-07-20 20:58:33  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  983  time:  2023-07-20 20:58:34  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  984  time:  2023-07-20 20:58:35  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  985  time:  2023-07-20 20:58:36  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  986  time:  2023-07-20 20:58:37  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  987  time:  2023-07-20 20:58:38  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  988  time:  2023-07-20 20:58:39  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  989  time:  2023-07-20 20:58:40  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  990  time:  2023-07-20 20:58:41  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  991  time:  2023-07-20 20:58:42  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  992  time:  2023-07-20 20:58:43  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  993  time:  2023-07-20 20:58:44  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  994  time:  2023-07-20 20:58:45  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  995  time:  2023-07-20 20:58:46  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  996  time:  2023-07-20 20:58:47  code:  200
http client select node "192.168.1.2:63842" for service "service-default-default-hello-world.svc-latest" 
Hello world one  id:  997  time:  2023-07-20 20:58:48  code:  200
http client select node "192.168.1.2:64140" for service "service-default-default-hello-world.svc-latest" 
Hello world two  id:  998  time:  2023-07-20 20:58:49  code:  200
http client select node "192.168.1.2:63853" for service "service-default-default-hello-world.svc-latest" 
Hello world three  id:  999  time:  2023-07-20 20:58:50  code:  200

Process finished with the exit code 0



```