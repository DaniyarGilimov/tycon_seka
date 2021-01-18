var names = [
    "Serikbek Mukhametkaliov",
    "Taras Alenov",
    "Kholmatzhon Maulenev",
    "Aidos Tairov",
    "Amir Akhmetev"
    ];
    
    
    
    var items = [1,2,5,7,8];
    var endings = ["@mail.ru", "@gmail.com","@skype.com", "@drive.com"];
    var alll = [];
    
    
    for (let i = 0; i < names.length; i++) {
    
    var defo = {
        "ID" : 0,
        "Name" : "Samat",
        "Surname" : "Sultanov",
        "Age" : 18,
        "Image" : "default",
        "Sex" : 1,
        "About" : "test",
        "JobTitle" : "CTO",
        "Phone" : "+7702153****",
        "Email" : "example@gmail.com",
        "Level" : 1,
        "Jobs" : [
                {
                        "StartYear" : 2016,
                        "EndYear" : 2018,
                        "Company" : "Jigi LLP",
                        "Location" : "Казахстан",
                        "Comment" : "С усердием работал, и показал все свои наилучшие стороны",
                        "Title" : "Обучение в сфере управления"
                },
                {
                        "StartYear" : 2019,
                        "EndYear" : 2020,
                        "Company" : "Google Inc",
                        "Location" : "США",
                        "Comment" : "Поднял инженерию на новый уровень",
                        "Title" : "CTO компании"
                }
        ],
        "Salary" : 5,
        "Price" : 5
    };
    
        let sp = names[i].split(" ");
        let z = i+1;
        let f = Math.floor(Math.random() * 9);
        let s = Math.floor(Math.random() * 9);
        let t = Math.floor(Math.random() * 9);
        let zer = items[Math.floor(Math.random() * items.length)];
        defo.ID = 75+i;
        defo.Phone = "+770"+zer+f+s+t+"****";
        defo.Name = sp[0];
        defo.Surname = sp[1];
        defo.Age = Math.floor(Math.random() * 15)+18;
        defo.Email = sp[0]+sp[1]+endings[Math.floor(Math.random() * items.length)];
        defo.Image = "https://azionline.kz:4041/workers/cto/men?name=LmpwZw ("+z+").jpg";
        defo.Level = Math.floor((i/5)+1);
        alll.push(defo);
    }
    console.log(JSON.stringify(alll));