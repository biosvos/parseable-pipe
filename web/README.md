문제. 토픽에 해당하는 목록만 보여줘야 한다.

1. 각 paragraph마다 class를 특정 클래스는 hidden 하게 한다.
2. 각 토픽마다 div를 따로 두고, 해당 div만을 hidden 한다.
3. 각 paragraph마다 토픽을 class로 두고 hidden 해야 하면, css를 추가해 히든한다.

선택 3
가장 쉬워보임

분석
- 동적으로 css 클래스를 둘수 있나?
- css 클래스간 우선순위는 어떻게 되나?

css 우선순위
1. 속성 값 뒤에 !important 를 붙인 속성
2. HTML에서 style을 직접 지정한 속성
3. #id 로 지정한 속성
4. .클래스, :추상클래스 로 지정한 속성
5. 태그이름 으로 지정한 속성
6. 상위 객체에 의해 상속된 속성

동적 css 추가
```javascript
var stylesheet = $("<link>", {
    rel: "stylesheet",
    type: "text/css",
    href: "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
});
stylesheet.appendTo("head");
```