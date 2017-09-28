# boyconf

Nuttycase: "the Best Configuration By JSON, For Golang Project"

I spent one day write those code, just want a better configure
lib: esay, humanreadble, support env, support autoload, and load trigger.

Now, It's running good.

## Install

`go get github.com/2guotou/boyconf`

## Config Struct

1. define a struct type for you configuartion
2. then make one config file

**Struct Sample**

```golang
type Conf struct{
    A bool                        //Cover
    B string                      //Cover
    C map[string]string           //Merge & Cover
    D map[string]struct{E string} //Merge & Cover
    F []string                    //Cover
}

var Cnf = new(Conf)

func main(){

    boy := &boyconf.Boy{
		File:       "path/to/test.conf",      //Config File Path
		Config:     Cnf,                      //The Global Config Variable
		Env:        []string{"product","en"}, //The Right Env Alway Has High Priority
		AutoReload: true,                     //Support Auto load, Default == false
	}

    trigger1 := func() {
		fmt.Println("I am trigger1, I will do sth.")
	}

    trigger2 := func() {
		fmt.Println("I am trigger2, I will run sth.")
	}

    err := boyconf.Init(boy, trigger1, trigger2)

    //will Print: &{A: false B: "" C:{c3: "ccc3"} D: map[d1:{e2:"eeee2" d2:{e3: "eee3"}}] F: ["f3"] }
	fmt.Printf("%+v, %v", Cnf, err)
	
    //You Can Change Your Config File, Reload Has 5 Seconds Latency
    //For Example: change product.D.d1.e2 = "eeeeeeeeeeeeeeeee2"
    time.Sleep(20 * time.Second)

    //will Print: &{A: false B: "" C:{c3: "ccc3"} D: map[d1:{e2:"eeeeeeeeeeeeeeeee2" d2:{e3: "eee3"}}] F: ["f3"] }
	fmt.Printf("%+v", Cnf)
}
```

**If Your Config File Look Like**

`test.conf`

```json
{
    "default" : {
        "A": true,
        "B": "bbb1",
        "C": {
            "c1":"ccc1",
            "c2":"ccc2"
        },
        "D": {
            "d1":{
                "e1": "eeee1"
            }
        },
        "F": ["f1","f2"]
    },
    "product":{
        "C": {
            "c3":"ccc3"
        },
        "D": {
            "d1":{
                "e2": "eeee2"
            },
            "d2":{
                "e3": "eeee3"
            }
        },
        "F": ["f3"]
    },
    "en": {
        "A": false
    },
    "jp":{
        "B": "bbb2"
    }
}
```
