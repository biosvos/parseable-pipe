문제. 토픽에 해당하는 목록만 보여줘야 한다.

<details>

발산.
1. 각 paragraph마다 class를 특정 클래스는 hidden 하게 한다.
2. 각 토픽마다 div를 따로 두고, 해당 div만을 hidden 한다.
3. 각 paragraph마다 토픽을 class로 두고 hidden 해야 하면, css를 추가해 히든한다.

수렴.
3번 가장 쉬워보임

분석.
- 동적으로 css 클래스를 둘수 있나?
- css 클래스간 우선순위는 어떻게 되나?

<details><summary>css 우선 순위</summary>

1. 속성 값 뒤에 !important 를 붙인 속성
2. HTML에서 style을 직접 지정한 속성
3. #id 로 지정한 속성
4. .클래스, :추상클래스 로 지정한 속성
5. 태그이름 으로 지정한 속성
6. 상위 객체에 의해 상속된 속성

</details>

<details><summary>동적 css 추가</summary>

```javascript
var stylesheet = $("<link>", {
    rel: "stylesheet",
    type: "text/css",
    href: "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
});
stylesheet.appendTo("head");
```

</details>

</details>

---
문제. 내용이 많으면 스크롤이 생기면서 아래 내용을 확인할 수 없다.
아래 내용을 확인하고 싶다.  
왜? 최신 내용이 가장 중요하다.

<details>

발산.
1. 최신 내용만 중요하다면, 기존 내용을 지워서 한페이지만 유지하자.
    - 기존 내용도 중요하다. 또한 최신 내용이 너무 길다면 최신 내용의 일부가 사라질 수 있다.
2. 스크롤을 맨 아래로 유지하자.
    - 언제? 새로운 내용이 추가될 때? 메뉴를 변경할 때?
3. 메뉴가 변경될 때, 스크롤을 자동으로 내리자.
새로운 내용이 추가될 때는 이미 스크롤이 맨 아래일 경우에만 계속 내리자.

수렴.
3번. 2번과 비교해 사용자가 멈춰있기를 원할 때 멈춰 있을 수 있다.

분석.
- 프로그래밍적으로 스크롤을 아래로 가게 할 수 있는가?
- 프로그래밍적으로 스크롤이 맨 아래에 있다는 것을 알 수 있는가?

<details><summary>가장 아래로 스크롤</summary>

```javascript
$("#mydiv").scrollTop($("#mydiv")[0].scrollHeight);
```b

</details>

<details><summary>스크롤이 가장 아래인지 확인</summary>

```javascript
$("#mydiv").scrollTop() === $("#mydiv")[0].scrollHeight;
```

</details>

</details>

---

문제. 활성화 되지 않은 탭에서 내용이 변경될 때, 탭을 클릭하지 않고도 탭의 내용이 변경되었다는 것을 알고 싶다.

<details>

발산.
1. 활성화되지 않는 탭에 새로운 내용이 추가되면 탭의 색을 변경한다. 
2. 활성화되지 않는 탭에 새로운 내용이 추가되면 탭의 내용 옆에 빨간 동그라미를 붙인다.
3. 활성화되지 않는 탭에 새로운 내용이 추가되면 해당 탭으로 전환한다.

수렴.
2번. 가장 직관적이고 쉽다.

</details>

---
문제. contents div 에 스크롤이 생기는게 아니라, html에 생김