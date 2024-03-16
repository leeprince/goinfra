<p> OAuth 2.0是一个应用之间彼此访问数据的开源授权协议。比如，一个游戏应用可以访问Facebook的用户数据或者一个基于地理的应用可以访问Foursquare的用户数据等。下面是一张阐述该概念的图： </p>

<p><img src='https://cloud.itgogogo.cn/typora/01、OAuth 2.0 引言/1654782667256.png'></p>

<blockquote> 
<p>OAuth 2.0怎么通过应用共享数据的例子</p>
 </blockquote>

<p> 用户访问web游戏应用，该游戏应用要求用户通过Facebook登录。用户登录到了Facebook,再重定向会游戏应用， 游戏应用就可以访问用户在Facebook的数据了，并且该应用可以代表用户向Facebook调用函数(如发送状态更新)。 </p>



##  OAuth 2.0实用案例
<p> OAuth 2.0要么用来创建一个能够从其他应用读取用户信息的应用(如上面图表中的游戏应用)，要么创建一个使其他应用访问自己的用户数据的应用(如上面例子中的Facebook)。OAuth 2.0是OAuth 1.0的替代品，OAuth 1.0更加复杂。OAuth 1.0涉及到了证书等，而OAuth 2.0更简单，它不需要任何证书，仅仅就SSL/TLS。 </p>



##  OAuth 2.0规范
<p> 该指南的目标是提供一个OAuth 2.0的很容易理解的概述，但是不会描述规范的每一个细节。如果你想实现OAuth 2.0, 你将很有可能要全面学习该规范 </p>

<p> 你可以在这里找到该规范：<a target="_blank" rel="nofllow" href="http://tools.ietf.org/html/draft-ietf-oauth-v2-23">http://tools.ietf.org/html/draft-ietf-oauth-v2-23</a> </p>

