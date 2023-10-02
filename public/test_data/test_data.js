let target1 = {
    targettext: "Разработка анти-чита для GTA ONLINE",
    done: "123400",
    target: "1000000",
    targettype: "долларов",
    doneproc: "10",
    leftproc: "90",
}

let target2 = {
    targettext: "Поддержка Red Dead Online",
    done: "100",
    target: "250000",
    targettype: "долларов",
    doneproc: "20",
    leftproc: "80",
}

let sublevel1 = {
    name: "Начинающий автоугонщик",
    price: "2000",
    description: "Ранний доступ к новым DLC",
}

let sublevel2 = {
    name: "Профи своего дела",
    price: "4000",
    description: "Предыдущий уровень + возможность предлагать идеи для обновлений",
}

let sublevel3 = {
    name: "Легенда улиц",
    price: "10000",
    description: "Предыдущий уровень + закрытый discord-сервер с разработчиками",
}

let sublevel4 = {
    name: "Контрольный пакет акций",
    price: "100000",
    description: "Предыдущий уровень + личная встреча с CEO",
}

let comment1 = {
    nickname: "VitoScaletta",
    text: "Лучше так, чем поторопиться и обрезать половину игры, как в Mafia II",
    datetime: "28 сен 2023 19:00",
    picname: "vito.jpg"
}

let comment2 = {
    nickname: "carlmarx",
    text: "Хватит вытягивать у людей деньги своим онлайном!!!!!!!! Уже скатились до уровня мобильных игр с их бесконечными микротранзакциями",
    datetime: "28 сен 2023 19:10",
    picname: "su.png"
}

let post1 = {
    availability: "$1000000",
    unlocked: false,
    datetime: "27 сен 2023 18:00",
    topic: "Процесс разработки GTA VI"
}

let post2 = {
    availability: "Доступен",
    unlocked: true,
    datetime: "26 сен 2023 18:00",
    topic: "Когда выйдет GTA VI?",
    text: "Нас очень часто спрашивают, когда же наконец выйдет GTA VI? Отвечаем: игра находится в разработке и пробудет там ещё не меньше 2 лет. Мы чувствуем большую ответственность перед своей аудиторией после ошеломительного успеха GTA V и GTA ONLINE, поэтому не хотим выпускать сырой продукт. А пока наслаждайтесь многочисленными обновлениями на GTA ONLINE!",
    tags: ["GTA VI", "gamedev"],
    comments: [comment1, comment2]
}

export let profile1 = {
    targets: [target1, target2],
    sub_levels: [sublevel1, sublevel2, sublevel3, sublevel4],
    posts: [post1, post2],
    subs: 229,
    nickname: "Rockstar Games",
    status: "Обновление 1.100 на GTA Online уже скоро!",
    about: "Маленькая инди-студия из Нью-Йорка. Известные проекты: Grand Theft Auto, Max Payne, Red Dead Redemption, The Bully, Midnight Club.",
    author: true
}

export let profile2 = {
    donated: 2229,
    nickname: "fanat 2chera",
    status: "=ТЕРПиМ КАРЛИКИ=",
    author: false,
    subs: [{name: "Морген"}, {name: "Морген2"}]
}