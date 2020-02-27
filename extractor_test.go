package textractor

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func Test_emptyNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			input: `<div></div>`,
			want:  false,
		},
		{
			input: `<h1></h1>`,
			want:  true,
		},
		{
			input: `<h1>test</h1>`,
			want:  false,
		},
		{
			input: `<section></section>`,
			want:  true,
		},
		{
			input: `<span></span>`,
			want:  true,
		},
		{
			input: `<span>        </span>`,
			want:  true,
		},
		{
			input: `<span>   
			     </span>`,
			want: true,
		},
		{
			input: `<span><b></b></span>`,
			want:  false,
		},
	}
	for _, tt := range tests {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(tt.input))
		assert.Nil(t, err, "")
		if got := canBeRemove(dom.Find("body").Children()); got != tt.want {
			t.Errorf("emptyNode() = %v, want %v, input \"%v\"", got, tt.want, tt.input)
		}
	}
}

var testSource = `
<!DOCTYPE html>
<html>

<head>
	<meta charset="UTF-8">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=0">
	<meta http-equiv="content-language" content="zh-CN">
	<link rel="shortcut icon" href="https://cdn.iyunhui.com/common/image/favicon.ico">	<title>10月16日云汇有色金属晚观察 - 铝业资讯</title>
	<meta name="keywords" content="云汇有色金属晚观察,铝价预测,铜价预测">
	<meta name="description" content="今日收盘后共采集到21位分析师观点：对于下周一铜价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点8位。对于下周一铝价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点7位。">
	<script src="https://cdn.iyunhui.com/lib/es6-shim-0.35.1/es6-shim.min.js"></script>
	<link rel="stylesheet" href="https://cdn.iyunhui.com/lib/layui-v2.3.0/css/layui.css">	<link rel="stylesheet" href="https://cdn.iyunhui.com/lib/icomoon/style.css?20190730">	<script src="https://cdn.iyunhui.com/lib/jquery-2.2.4/jquery.min.js"></script>	<link rel="stylesheet" href="https://cdn.iyunhui.com/common/common.css?20191108">	<script src="https://cdn.iyunhui.com/lib/moment-2.20.1/moment.min.js"></script>
<script src="https://cdn.iyunhui.com/lib/moment-2.20.1/locale/zh-cn.js"></script>	<script src="https://cdn.iyunhui.com/lib/lodash-4.17.5/lodash.min.js"></script>	<script src="https://cdn.iyunhui.com/lib/axios-0.17.1/axios.min.js"></script>
<script src="https://cdn.iyunhui.com/lib/blueimp-md5-2.10.0/md5.min.js"></script>
<script src="https://cdn.iyunhui.com/common/sdk-1.0.0/sdk.js?20200114"></script>	<link href="https://cdn.iyunhui.com/passport/index-login.css?20200114" rel="stylesheet">	<script src="https://cdn.iyunhui.com/lib/jwt-decode/jwt-decode.min.js"></script>
<script src="https://cdn.iyunhui.com/passport/sso.js?2020011401"></script>		<script src="//msite.baidu.com/sdk/c.js?appid=1622257158322537"></script>
    <script type="application/ld+json">
        {
            "@context": "https://ziyuan.baidu.com/contexts/cambrian.jsonld",
            "@id": "https://al.iyunhui.com/news/151456/",
            "appid": "1622257158322537",
			"title": "10月16日云汇有色金属晚观察 - 铝业资讯",
			            "pubDate": "2019-10-16T21:24:04",
            "upDate": "2019-10-16T21:24:04"
        }
	</script>
		<link href="https://cdn.iyunhui.com/common/market-news.css?20191231" rel="stylesheet">
	<link href="https://cdn.iyunhui.com/news/index.css?20190801" rel="stylesheet">
			<meta property="og:type" content="article"/>
	<meta property="article:published_time" content="2019-10-16T21:24:04" /> 
	<meta property="og:title" content="10月16日云汇有色金属晚观察" />
	<meta property="og:description" content="今日收盘后共采集到21位分析师观点：对于下周一铜价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点8位。对于下周一铝价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点7位。" />
	<meta property="og:url" content="https://al.iyunhui.com/news/151456/" />
		<meta property="article: author" content="铝云汇" />
	<link rel="alternate" media="only screen and(max-width: 640px)" href="https://5g.iyunhui.com/al/news-151456/" >
	<meta http-equiv="mobile-agent" content="format=html5; url=https://5g.iyunhui.com/al/news-151456/" />
	<link rel="miphtml" href="https://m.iyunhui.com/al/news/151456/">
	</head>

<body><div class="layui-header yh-header">
	<div class="layui-container">
		<ul class="layui-nav yh-nav pull-left">
			<li class="layui-nav-item  yh-home-list">
				<a href="/"><i class="iconfont icon-index-1 font19"></i>铝云汇</a>
				<div class="list">
										<div class="this">
						<a href="//al.iyunhui.com">铝云汇</a>
					</div>
					<div class="other">
						<a href="//cu.iyunhui.com" target="_blank">铜云汇</a>
					</div>
										<div class="other">
						<a href="//www.iyunhui.com" target="_blank">九商云汇</a>
					</div>
				</div>
			</li>
			<li class="layui-nav-item ">
				<a href="/product/"><i class="iconfont icon-product font20"></i>产品</a>
			</li>
			<li class="layui-nav-item ">
				<a href="/company/"><i class="iconfont icon-shop font20"></i>公司</a>
			</li>
			<li class="layui-nav-item layui-this">
				<a href="/news/"><i class="iconfont icon-news font20"></i>资讯</a>
			</li>
			<li class="layui-nav-item ">
				<a href="/market/"><i class="iconfont icon-chart-line font20"></i>行情</a>
			</li>
			<li class="layui-nav-item layui-nav-item-51 ">
				<a href="https://www.iyunhui.com/feed/"><i class="yh-iconfont-new"></i>快消息</a>
			</li>
			<li class="layui-nav-item layui-nav-item-51 ">
				<a href="https://www.iyunhui.com/price/"><i class="yh-iconfont-price"></i>有色价格</a>
			</li>
			<li class="layui-nav-item">
								<a href="https://data.iyunhui.com" target="_blank"><i class="iconfont icon-data font19"></i>数据中心</a>
							</li>
		</ul>
		<div class="pull-right margin-top-17 layui-inline">
	<span class="color-ccc" style="font-family: 宋体;">
	<span class="em1"></span>/
	<span class="em0"></span>
		</span><a href="https://bigal.iyunhui.com/news/151456/">繁体</a>
	</div>		<span id="login-top-bar">
			<div class="pull-right margin-top-17 layui-inline">
				<a href="https://passport.iyunhui.com" class="color-888">登录</a>
				<span class="color-ccc" style="font-family: 宋体;">
					<span class="em0"></span>/
					<span class="em0"></span>
				</span>
				<a href="https://passport.iyunhui.com/?a=reg" class="color-888">注册</a>
			</div>
		</span>
	</div>
</div>
<script>
$(function(){
	checkLogin();
})
</script><div class="yh-containter-header">
	<div class="layui-container">
		<div class="logo logo-img logo-img-al">
			<a href="/">
				<h2 class="layui-hide">铝云汇</h2>
			</a>
			<div class="yh-logo-title"><span class="layui-hide">铝业</span>资讯</div>
		</div>
		<div class="pull-right">
			<div class="pull-right margin-top-25 margin-left-15 layui-inline yh-qrcode">
				<span class="font12 pull-right w2">手机
					<br />云汇</span>
				<i class="icon-scan-qrcode font44 color-999 pull-right"></i>
				<div class="sn-qrcode">
					<div class="sn-qrcode-content">
						<img src="https://cdn.iyunhui.com/common/image/sn-qrcode-al.png" />
					</div>
					<p>扫一扫，关注手机资讯</p>
					<b></b>
				</div>
			</div>
			<div class="yh-header-search">
				<input type="text" placeholder="请输入关键字" value="" class="layui-input pull-left search-text">
				<a href="javascript:;" class="pull-right yh-btn search-btn">
					<i class="icon-search font19"></i>
				</a>
			</div>
			<ul class="layui-nav yh-nav2">
				<li class="layui-nav-item">
                    <a href="/news/">首页</a>
                </li>
                <li class="layui-nav-item">
                    <a href="/news/t-29/">行业形势</a>
                </li>
                <li class="layui-nav-item ">
                    <a href="/news/t-30/">企业动态</a>
                </li>
                <li class="layui-nav-item ">
                    <a href="/news/t-34/">近期热点</a>
                </li>
                <li class="layui-nav-item">
                    <a href="/news/t-36/">权威解读</a>
                </li>
			</ul>
		</div>
	</div>
</div>
<div class="layui-container margin-top-20">
    <span class="layui-breadcrumb" lay-separator="&gt;" style="visibility: visible;">
        <a href="/">铝云汇</a>
        <a href="/news/"><span class="layui-hide">铝业</span>资讯</a>
        <a href="https://al.iyunhui.com/news/151456/">
            <cite>正文</cite>
        </a>
    </span>
</div>

<div class="layui-container yh-detail-mod">
    <h1 class="mod-header">10月16日云汇有色金属晚观察</h1>
    <div class="mod-header2">
        <i class="icon-time-1 font16 color-888"></i>
        <span class="color-888">2019年10月16日 21:24:04</span>
        <span class="em1"></span>
                铝云汇
                <span class="em1"></span>        <span class="em1"></span>
        <div class="rating">
                        <i class="icon-rate wb-star orange-600"></i>
                        <i class="icon-rate wb-star orange-600"></i>
                        <i class="icon-rate wb-star orange-600"></i>
                        <i class="icon-rate wb-star orange-600"></i>
                        <i class="icon-rate wb-star orange-600"></i>
                    </div>
        <div class="layui-inline pull-right relate-keyword">
                        <a href="/news/k-6pJ8sq86q9Ysb2X7mE83yau97kz-1b10c1/" target="_blank">云汇有色金属晚观察</a>
                        <a href="/news/k-bfK6rpbv58wG-517a7c/" target="_blank">铝价预测</a>
                        <a href="/news/k-bfJ6rpbv58wG-d0439a/" target="_blank">铜价预测</a>
                    </div>
    </div>
    <div class="mod-content row">
        <div class="article-left-btn-group">
	<ul class="share">
		<li>
			<a href="https://service.weibo.com/share/share.php?url=https://al.iyunhui.com/news/151456/&title=10月16日云汇有色金属晚观察&appkey=" class="icon-article icon-weibo font40" target="_blank"></a>
		</li>
		<li class="weixin">
			<a href="#" class="icon-article icon-weixin font40"></a>
			<div class="yh-qrcode-tc">
				<b></b>
				<div class="share-qrcode"></div>
			</div>
		</li>
		<li>
			<a href="http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=https://al.iyunhui.com/news/151456/&title=10月16日云汇有色金属晚观察&desc=今日收盘后共采集到21位分析师观点：对于下周一铜价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点8位。对于下周一铝价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点7位。&summary=今日收盘后共采集到21位分析师观点：对于下周一铜价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点8位。对于下周一铝价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点7位。&site=九商云汇" class="icon-article icon-qzone font40" target="_blank"></a>
		</li>
		<li class="tel">
			<a href="#" class="icon-article icon-qrcode font40"></a>
			<div class="yh-qrcode-tc">
				<b></b>
				<div class="share-qrcode"></div>
			</div>
		</li>
	</ul>
</div>
<script src="https://cdn.iyunhui.com/lib/jquery-qrcode-1.0/jquery.qrcode.min.js"></script>
<script>
$(".share-qrcode").qrcode({width: 160,height: 160,text: "https://al.iyunhui.com/news/151456/"});
</script>        <div class="layui-col-xs9 yh-col-xs9-900">
            <div class="market-detail">
                <p style="text-indent: 2em; text-align: left;">各位，久等，今天是10月16日，星期三。</p><p style="text-indent: 2em; text-align: left;">首先我们来回顾今日有色行业要闻：</p><p style="text-indent: 2em; text-align: left;">1、9月份铜板带箔企业开工率为70.62%，同比下滑6.16个百分点，环比回升1.16个百分点。</p><p style="text-indent: 2em; text-align: left;">2、南山铝业将于2020年底在印尼投产一期氧化铝项目，第一阶段项目建设已从今年年初开始，具有同样产能的第二阶段项目预计将于明年年中启动建设。统计显示，明年年底前中国公司在印尼投资的的氧化铝项目合计年产能将达400万吨。</p><p style="text-indent: 2em; text-align: left;">3、由于社区抗议、封锁道路，切断了物资供应，五矿资源已暂停了Las Bambas铜矿90%的运营，周三或将完全停止该矿运营。秘鲁政府试图重启与当地社区的对话。</p><p style="text-indent: 2em; text-align: left;">4、泰克资源旗下智利铜矿罢工进入第二天，该矿工会主席Manuel Alvarez表示，劳资双方未安排任何谈话。</p><p style="text-indent: 2em; text-align: left;">5、印尼金川红土镍矿冶炼厂1#线镍铁电炉砌筑施工完成，标志着印尼金川红土镍矿项目所有耐材施工全部结束。</p><p style="text-indent: 2em; text-align: left;">6、CRU的估计数据，今年前三季度世界精铜产量为1765.10万吨, 消费量为1769.50万吨，短缺数量为4.40万吨。</p><p style="text-indent: 2em; text-align: left;">7、智利安托法加斯塔矿业公司周二表示，已与智利洛斯佩兰贝尔斯旗舰铜矿的工会达成劳工协议。</p><p style="text-indent: 2em; text-align: left;">8、惠誉解决方案周三预计，中国2019-2028年铜产量年增长速度平均奖放缓至1.9%，低于前10年平均4.8%的增速，因国内铜矿石品位降低。</p><p style="text-indent: 2em; text-align: left;">9、哈萨克斯坦统计局公布的数据显示，今年1-9月精炼铜产量同比增加11%，至354,836吨；精炼锌产量同比下滑0.9%，至235,914吨。</p><p style="text-indent: 2em; text-align: left;">10、力拓公布的报告显示，2019年第三季度铝土矿产量为1380万吨，同比增加9%。</p><p style="text-indent: 2em; text-align: left;">11、韩国央行将基准利率下调25个基点至1.25%，为年内第二次降息，追平历史低点。</p><p style="text-indent: 2em; text-align: left;">关键性指数变动情况：</p><p style="text-align:center"><img src="https://img.iyunhui.com/news/20191016/d502d74eecb26c4c077180174dd61547.png/watermark"/></p><p style="text-indent: 2em; text-align: left;">美元指数(USDX)收盘报98.30点，较上一交易日跌0.16点，跌幅0.16%；CRB指数收盘报174.75点，较上个交易日跌0.24点，跌幅0.14%；美国原油(WTI)主力合约收盘报52.91美元/桶，较上个交易日跌0.57美元，跌幅1.07%；BDI(波罗的海干散货运价指数)较上个交易日1898.00点，收报跌18点，跌幅0.94%。上证综指收盘报2978.71点，较上个交易日跌12.34点，跌幅0.41%。</p><p style="text-indent: 2em; text-align: left;">【铝】</p><p style="text-indent: 2em; text-align: left;">期货方面：</p><p style="text-indent: 2em; text-align: left;">沪铝1911收报13885涨120，沪铝2001收报13830涨70；LME场内铝收报1728涨8。</p><p style="text-indent: 2em; text-align: left;">现货方面：</p><p style="text-indent: 2em; text-align: left;">今日期铝涨幅较大，前期惜售的持货商今日出货积极性有明显提升，而中间商接货热情未减，部分贸易商认为现货市场偏紧，对短期内铝价及升水有一定期待，因而接货较为积极，买卖双方交投活跃。因铝价大涨，部分下游今日按需走货为主，但也有贸易商反馈少量厂商因前期对价格观望一直买货不多，如今库存消耗无几不得不开始接货。整体而言，华东今日整体成交不错。</p><p style="text-indent: 2em; text-align: left;">价格方面：</p><p style="text-indent: 2em; text-align: left;">长江有色A00铝均价报13880涨110；南海有色佛山A00铝均价报14150涨70；中铝公司华东铝锭13900，华南铝锭13860，西南铝锭13820，中原铝锭13800，氧化铝地区报价全线上调50元/吨。</p><p style="text-indent: 2em; text-align: left;">关键性指数：</p><p style="text-indent: 2em; text-align: left;">上海期货、长江有色基差报5，较前一交易日跌5；沪伦比报7.97，较前一交易日跌0.25%；升贴水报-55，较前一交易日跌40；长江有色、南海有色南北价差报270，较前一交易日跌40；铝价指数报93.13，较前一交易日涨0.64或涨幅0.69%。</p><p style="text-indent: 2em; text-align: left;">【铜：沪伦比触发预警】</p><p style="text-indent: 2em; text-align: left;">期货方面：</p><p style="text-indent: 2em; text-align: left;">沪铜1911收报46730跌180，沪铜2001收报46800跌160；LME场内铜收报5773跌45；COMEX期铜主力合约收跌0.02美元，收盘报2.6115美元/磅。</p><p style="text-indent: 2em; text-align: left;">现货方面：</p><p style="text-indent: 2em; text-align: left;">换月后首日，持货商持稳报价于升水40-升水70元/吨，低价货源得到市场青睐，吸引贸易商入市收货，早市尚能压价，成交活跃度尚好，尤其是平水铜买兴较高，询盘明显较为活跃，但是周初进口铜集中到货入库令库存压力仍在，成交暂难有实质性改变，下游按需采购。</p><p style="text-indent: 2em; text-align: left;">价格方面：</p><p style="text-indent: 2em; text-align: left;">长江有色1#铜均价报46960跌120；南储华南1#阴极铜均价报46880跌110。</p><p style="text-indent: 2em; text-align: left;">关键性指数：</p><p style="text-indent: 2em; text-align: left;">上海期货、长江有色基差报-230，较前一交易日跌30；沪伦比报8.13，较前一交易日涨0.74%，连续两天触发预警；升贴水报70，较前一交易日跌10；长江有色、广东南储南北价差报-80，较前一交易日涨10；铜价指数报95.3，较前一交易日跌0.24或跌幅0.25%。</p><p style="text-indent: 2em; text-align: left;">【铅锌镍锡】</p><p style="text-indent: 2em; text-align: left;">期货方面：</p><p style="text-indent: 2em; text-align: left;">沪锌1911收报18880跌100，沪锌2001收报18955跌40；沪铅1911收报16895涨75，沪铅2001收报16675涨60；沪镍1911收报133820跌950，沪镍2001收报131550跌1480；沪锡1911收报136040涨140，沪锡2001收报138790涨1830。</p><p style="text-indent: 2em; text-align: left;">LME场内锌收报2438涨18；LME场内铅收报2148.50涨17.5；LME场内镍收报16995涨345；LME场内锡收报16700涨180。</p><p style="text-indent: 2em; text-align: left;">价格方面：</p><p style="text-indent: 2em; text-align: left;">长江有色1#铅均价报16925涨50；长江有色0#锌均价报19480跌100；长江有色1#锌均价报18955跌100；长江有色1#镍板均价报134600涨300；长江有色1#锡均价报138750涨1250。</p><p style="text-indent: 2em; text-align: left;">【有色金属库存情况】</p><p style="text-align:center"><img src="https://img.iyunhui.com/news/20191016/e5f946e3660c41323bcf2f13a8a42310.png/watermark"/></p><p style="text-align:center"><img src="https://img.iyunhui.com/news/20191016/155e709b4b4ef876d555a15ae98074ec.png/watermark"/></p><p style="text-indent: 2em; text-align: left;">上期所基本金属仓单：铜64399吨，增加1472吨；铝98686吨，增加3178吨；锌35659吨，减少153吨；铅14710吨，持平；镍23757吨，持平；锡3965吨，增加84吨。</p><p style="text-indent: 2em; text-align: left;">伦敦金属交易所（LME）：LME铜库存增加600吨，铝库存减少2375吨，锌库存减少150吨，镍库存减少3600吨，锡库存增加130吨，铅库存持平。</p><p style="text-indent: 2em; text-align: left;">COMEX铜库存报37408短吨，较上个交易日减22短吨。</p><p style="text-indent: 2em; text-align: left;">【分析汇总】</p><p style="text-indent: 2em; text-align: left;">今日收盘后共采集到21位分析师观点：</p><p style="text-indent: 2em; text-align: left;">对于下周一铜价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点8位。</p><p style="text-indent: 2em; text-align: left;">对于下周一铝价持看涨及震荡偏强观点3位，持看跌及震荡偏弱观点7位。</p><p style="text-indent: 2em; text-align: left;">今日的晚观察就是这样，我们明天见！</p>            </div>
            <div class="yh-detail-qnode-two">
                <img src="https://cdn.iyunhui.com/common/image/snode3.jpg">
                <a class="text-con" href="https://data.iyunhui.com">有色数据中心：https://data.iyunhui.com</a>
            </div>
            <div class="yh-pre-new">
                上一篇：
                <a href="/news/151407/" class="font24 bst">贵阳市委书记赵德明到中铝股份贵州分公司调研</a>
            </div>
            <div class="yh-statement margin-top-20">未经许可，不得转载或以其他方式使用本网的原创内容。来源未注明或非铝云汇的文章，刊载此文出于传递更多信息之目的，不代表本网赞同其观点和对其真实性负责。所有文章内容仅供参考，不构成决策建议。</div>
<div class="margin-top-20"></div>            <div class="yh-market-news margin-top-40">
                <div class="yh-market-news-tit margin-top-10">
                    <h4 class="yh-mod-header inline-block">资讯</h4>
                </div>
                <div class="yh-grid-news-list yh-grid-news-list2">
                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="ceshi" href="/news/151712/">ceshi</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-100-d2c558/" target="_blank" class="font14 color-999"></a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-23 09:46:32</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>4</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess"></div>
	</div>
</div>                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="“退城进园”，贵州华仁打造千亿级生态循环铝工业基地" href="/news/151622/">“退城进园”，贵州华仁打造千亿级生态循环铝工业基地</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-6JC6qw7ZL87A7Zo-d3b03c/" target="_blank" class="font14 color-999">华仁新材料</a>
								<span class="em1"></span>
												<a href="/news/k-bfK7t46nM-5d4244/" target="_blank" class="font14 color-999">铝工业</a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-17 09:16:16</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>11</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess">作为西南地区发展铝工业优势明显的城市，贵阳到2020年，将初步打造形成千亿级铝及铝加工产业集群。过去三年，省、市主要领导曾多次去到位于清镇市的贵州华仁新材料有限公司考察“贵州轻合金新材料退城进园”项目，这个项目为什么引起省、市领导如此高的关注？它在贵州铝工业发展中发挥着怎样的作用？本期《高标准要求高水平开放高质量发展》专栏，让我们一起走进贵州华仁新材料。　　近日，记者走进位于清镇市的千亿级生态循环</div>
	</div>
</div>                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="中铝西南铝前三季度成绩优异" href="/news/151617/">中铝西南铝前三季度成绩优异</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-6o6bfKat36JLbfK-4ce5e9/" target="_blank" class="font14 color-999">中铝西南铝</a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-17 09:12:00</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>11</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess">截至9月30日，中铝西南铝业（集团）有限责任公司今年前3季度完成商品产量同比增长11.5%，其中，重点产品产销两旺，同比增长26%，以实际行动谱写了“不忘初心、牢记使命”的新篇章。 高效生产保市场 “只有为客户不断提供更优质的产品和服务，企业才会赢得客户的信任，才能在不断的竞争中提升核心竞争能力和盈利能力。”西南铝党委书记、执行董事尹雪春今年以来多次</div>
	</div>
</div>                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="内蒙古打破高端铝产业链技术瓶颈" href="/news/151599/">内蒙古打破高端铝产业链技术瓶颈</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-6o67NKbfK6nM-0ebde6/" target="_blank" class="font14 color-999">中拓铝业</a>
								<span class="em1"></span>
												<a href="/news/k-94T7iebfK9vD-df4748/" target="_blank" class="font14 color-999">电子铝箔</a>
								<span class="em1"></span>
												<a href="/news/k-6B99vD-48a324/" target="_blank" class="font14 color-999">光箔</a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-17 09:00:24</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>33</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess">记者16日从内蒙古自治区科技厅了解到，由内蒙古中拓铝业研发团队研制的自主知识产权电子铝箔光箔生产线实现多批量生产。产品投放市场后，用户采用后认为完全可以替代进口产品。这标志着内蒙古具有高附加值的电子铝箔产业链瓶颈终于被打破。　　据介绍，电子铝箔行业的产业链，首先是从铝锭中提炼出高纯铝，高纯铝通过熔铸、热轧、冷轧及热处理等工序加工成为电子铝箔光箔，光箔经腐蚀化成生产成用于铝电解电容器的电极箔。铝产业</div>
	</div>
</div>                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="美国铝业3Q净亏损2.21美元 宣布审查精炼战略以降低成本" href="/news/151542/">美国铝业3Q净亏损2.21美元 宣布审查精炼战略以降低成本</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-9MkbfK-96cc9d/" target="_blank" class="font14 color-999">美铝</a>
								<span class="em1"></span>
												<a href="/news/k-7ix7MZ-1b150b/" target="_blank" class="font14 color-999">季报</a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-17 08:43:21</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>3</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess">10月16日，全球铝土矿、氧化铝和铝产品的领导者美国铝业(AA.US)公布2019年第三季度业绩,期内实现收入26亿美元，净亏损2.21亿美元，或每股1.19美元;而2019年第二季度净亏损4.02亿美元，每股2.17美元。不计特殊项目，调整后的净亏损为8200万美元，或每股0.44美元。　　扣除特殊项目后的调整后EBITDA为3.88亿美元，比上一季度减少了6700万美元，这主要是由于氧化铝价格</div>
	</div>
</div>                                        <div class="yh-item-new">
	<div class="text-list">
				<h3>
						<a target="_blank" title="10月16日LME基本金属综述" href="/news/151469/">10月16日LME基本金属综述</a>
					</h3>
		<div class="time">
			<span>
												<a href="/news/k-2g02h029074e86Zb2X7mE-b7e455/" target="_blank" class="font14 color-999">LME基本金属</a>
								<span class="em1"></span>
							</span>
			<span class="em2"></span>
			<span>2019-10-17 07:06:06</span>
			<span class="em2"></span>
			<!-- <span>
				<i class="icon-eye font16"></i>
				<span class="em0"></span>2</span> -->
			<span class="em2"></span>
			<div class="rating">
								<i class="icon-rate wb-star orange-600"></i>
								<i class="icon-rate wb-star orange-600"></i>
							</div>
		</div>
		<div class="mess">伦敦10月16日消息，伦敦金属交易所(LME)期铜周三跌至一周低点，因中美贸易争端没有缓解的迹象，投资者越来越担忧这将拖累全球经济体质，给金属需求带来压力。伦敦时间10月16日17:00(北京时间10月17日00:00)，LME三个月期铜收跌0.7%，报每吨5,730美元，前一交易日跌幅与此相近。国际货币基金组织(IMF)表示，中美贸易摩擦拖累2019年全球经济增速降至2008-09年金融危机以来</div>
	</div>
</div>                                    </div>
            </div>
        </div>
        <div class="layui-col-xs3 yh-col-xs3-240 margin-top-40">
            <div class="yh-mod-hotmarket">
                <div class="mod-tit">
                    <div class="w1"></div>
                    <div class="w2">相关资讯</div>
                </div>
                <div class="mod-con">
                    <ul class="yh-grid-list">
                                                <li>
                                                        <a title="10月15日云汇有色金属晚观察" href="/news/150942/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191015/o_1dn7otfchs0b1bj81sta1de0106c7.jpg">
                            </a>
                                                        <a title="10月15日云汇有色金属晚观察" href="/news/150942/" class="font14 yh-text" target="_blank">
                                10月15日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="10月14日云汇有色金属晚观察" href="/news/150305/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191014/o_1dn58951a1cem1gc91gs713vdaps7.jpg">
                            </a>
                                                        <a title="10月14日云汇有色金属晚观察" href="/news/150305/" class="font14 yh-text" target="_blank">
                                10月14日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="铝云汇周观察（10.8-10.11）" href="/news/149962/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191012/o_1dmvt6rf810j8f2f8up1bo815k57.jpg">
                            </a>
                                                        <a title="铝云汇周观察（10.8-10.11）" href="/news/149962/" class="font14 yh-text" target="_blank">
                                铝云汇周观察（10.8-10.11）                            </a>
                        </li>
                                                <li>
                                                        <a title="10月11日云汇有色金属晚观察" href="/news/149761/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191011/o_1dmtc89on10os6ods7a1sog1ud97.jpg">
                            </a>
                                                        <a title="10月11日云汇有色金属晚观察" href="/news/149761/" class="font14 yh-text" target="_blank">
                                10月11日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="10月10日云汇有色金属晚观察" href="/news/149224/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191010/o_1dmqun7q3173q4bm1mem160h1e9j7.jpg">
                            </a>
                                                        <a title="10月10日云汇有色金属晚观察" href="/news/149224/" class="font14 yh-text" target="_blank">
                                10月10日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="10月9日云汇有色金属晚观察" href="/news/148557/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191009/o_1dmobsptf14o8bdbqae1vsj13b7.jpg">
                            </a>
                                                        <a title="10月9日云汇有色金属晚观察" href="/news/148557/" class="font14 yh-text" target="_blank">
                                10月9日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="10月8日云汇有色金属晚观察" href="/news/148043/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20191008/o_1dml9gfud19qh11n0g687101td7.jpg">
                            </a>
                                                        <a title="10月8日云汇有色金属晚观察" href="/news/148043/" class="font14 yh-text" target="_blank">
                                10月8日云汇有色金属晚观察                            </a>
                        </li>
                                                <li>
                                                        <a title="9月30日云汇有色金属晚观察" href="/news/146849/" target="_blank">
                                <img src="https://img.iyunhui.com/news/20190930/o_1dm11g27410g416021dkq1uqm7oc7.jpg">
                            </a>
                                                        <a title="9月30日云汇有色金属晚观察" href="/news/146849/" class="font14 yh-text" target="_blank">
                                9月30日云汇有色金属晚观察                            </a>
                        </li>
                                            </ul>
                </div>
            </div>
        </div>
    </div>
</div>
<script src="https://t.iyunhui.com/?term=zpjopa&type=news"></script><div class="yh-mod-link">
	<div class="yh-content">
		<div class="layui-container">
			<strong class="font20">媒体合作</strong>
		</div>
	</div>
	<div class="layui-container">
		<ul class="yh-grid-link">
															<li>
				<a href="http://www.cnmn.com.cn/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0j4i71ib61skm1coik6no377.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.99qh.com/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0just19di1pnn14ushf87rm7.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.baiinfo.com/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0knpq1fhgklk1vqv1t2g14427.jpg" />
				</a>
			</li>
												<li>
				<a href="http://nm.sci99.com/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0lcrklemf5v13131qkv1h187.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.qhrb.com.cn/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0m4s91o3chv9hvo4e5957.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.caijing.com.cn/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0mr421nks154mlms1u9t7bi7.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.f139.com/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0nvc51ihau9iha61ppo487.jpg" />
				</a>
			</li>
												<li>
				<a href="http://futures.hexun.com" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0rurp17f91c6m19t21b61oa97.jpg" />
				</a>
			</li>
												<li>
				<a href="http://www.metalsinfo.com/" target="_blank">
					<img src="https://img.iyunhui.com/link/o_1cim0scvn18gv1the1c2q2241bhk7.jpg" />
				</a>
			</li>
								</ul>
		<ul class="yh-link-text">
									<li>
				<a href="http://finance.qq.com/" target="_blank">
					腾讯财经				</a>
			</li>
																																																															<li><a href="//www.iyunhui.com/about/link.html" target="_blank">更多&gt;&gt;</a></li>
		</ul>
	</div>
</div>
<div class="yh-mod-footer">
    <div class="layui-container">
        <div class="layui-row">
            <div class="layui-col-xs4 width-l">
                <div class="font16 bst">关于我们</div>
                <ul class="yh-grid-about">
                    <li><a href="https://www.iyunhui.com/about/map.html" target="_blank">网站地图</a></li>
                    <li><a href="https://www.iyunhui.com/about/protocol.html" target="_blank">注册协议</a></li>
                    <li><a href="https://www.iyunhui.com/about/contact.html" target="_blank">联系我们</a></li>
                    <li><a href="https://www.iyunhui.com/about/money.html" target="_blank">汇款方式</a></li>
                    <li><a href="https://www.iyunhui.com/about/link.html" target="_blank">友情连接</a></li>
                </ul>
            </div>
            <div class="layui-col-xs4 width-c">
                <div class="font16 bst" style="margin-left: 100px;">汇服务</div>
                <ul class="yh-grid-about1">
                    <li><a href="//al.iyunhui.com" target="_blank">铝云汇</a></li>
                    <li><a href="//cu.iyunhui.com" target="_blank">铜云汇</a></li>
                    <li><a href="https://passport.iyunhui.com/?a=reg" target="_blank">会员注册</a></li>
                    <li><a href="https://passport.iyunhui.com/?a=login" target="_blank">会员登录</a></li>
                </ul>
            </div>
            <div class="layui-col-xs4 width-r">
                <div class="yh-icon-telephone"><i class="iconfont icon-tel"></i></div>
                <div class="yh-right-con">
                    <div class="w1">0371-61877310</div>
                    <div class="w2">周一至周五 9:00-17:00</div>
                    <div class="w3">
                        <a class="layui-btn layui-btn-danger" href="http://wpa.qq.com/msgrd?v=3&uin=3529191533&site=qq&menu=yes" target="_blank">
                            <i class="iconfont icon-qq"></i>在线客服
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="copy">
    Copyright&nbsp;&nbsp;©&nbsp;&nbsp;2020&nbsp;&nbsp;九商云汇&nbsp;&nbsp;版权所有<span class="em3"></span>豫ICP备18038776号
</div></div>

<div class="sideicon" style="display: block;">
    <div class="top backup" style="display: none;">
        <i class="icon-article icon-arrow-up-1"></i>
    </div>
    <a class="qq" href="http://wpa.qq.com/msgrd?v=3&uin=3529191533&site=qq&menu=yes" target="_blank">
        <i class="icon-article icon-qq"></i>
        <p style="display: none;">点击咨询</p>
    </a>
    <a class="phone" href="tel:075582484966">
        <i class="icon-article icon-phone"></i>
        <p style="display: none;">0371-61877310</p>
    </a>
    <div class="code">
        <i class="icon-article icon-qrcode-1"></i>
        <div class="codeimg3">
            <div class="share-qrcode">
                <img src="https://cdn.iyunhui.com/common/image/mp.png">
            </div>
            <b></b>
            <div>微信扫一扫</div>
        </div>
    </div>
</div>
<script>
$(function(){
    $(".sideicon .qq,.phone").mouseenter(function () {
        $(this).find("p").show();
    }).mouseleave(function () {
        $(this).find("p").hide();
    });
    $(window).scroll(function () {
        var top = $(window).scrollTop() - 350;
        if (top > 0) {
            $(".backup").show();
        } else {
            $(".backup").hide();
        }
    });
    $(".backup").click(function () {
        if ($('html').scrollTop()) {
            $('html').animate({ scrollTop: 0 }, 300);
            return false;
        }
        $('body').animate({ scrollTop: 0 }, 300);
        return false;
    });
});
</script>
<script>
    var _hmt = _hmt || [];
    (function() {
      var hm = document.createElement("script");
      hm.src = "https://hm.baidu.com/hm.js?97d440c597b3a48423a1aea4ab7372a4";
      var s = document.getElementsByTagName("script")[0]; 
      s.parentNode.insertBefore(hm, s);
    })();
</script>
<script>
    (function(){
    var src = (document.location.protocol == "http:") ? "http://js.passport.qihucdn.com/11.0.1.js?1c9f96153303bca4ce51175675ab233a":"https://jspassport.ssl.qhimg.com/11.0.1.js?1c9f96153303bca4ce51175675ab233a";
    document.write('<script src="' + src + '" id="sozz"><\/script>');
    })();
</script>
<script src="https://cdn.iyunhui.com/common/tj.js"></script>
<script>
    (function(){
        var bp = document.createElement('script');
        var curProtocol = window.location.protocol.split(':')[0];
        if (curProtocol === 'https') {
            bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
        }
        else {
            bp.src = 'http://push.zhanzhang.baidu.com/push.js';
        }
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(bp, s);
    })();
</script><script type="text/javascript" src="https://js.users.51.la/20190781.js"></script>
<script src="https://cdn.iyunhui.com/lib/layui-v2.3.0/layui.all.js"></script>
<script src="https://cdn.iyunhui.com/lib/current-device-0.8.0/current-device.min.js"></script>
<script>
    $(function(){
        if(device.type == "mobile"){
            var mipurl = "https://5g.iyunhui.com/al/news-151456/";
            var html = '<div class="switch-grid-fix"><div class="switch-grid">切换到<a data-type="mip" data-title="" href="'+mipurl+'" class="btn-tel">移动版</a></div></div>';
            $("body").append(html);
        }
    });
</script></body>

</html><script>
$(function(){
	$(".search-btn").click(function(){
		var str = $(this).parent().find(".search-text").val();
		toSearch(str);
		return false;
	});
	$(".search-text").keyup(function(e){
		if (e.keyCode == 13) {
			var str = $(this).val();
			toSearch(str);
		}
		return false;
	});
	function toSearch(str){
		var sdk = new iyhSDK();
		if(str != ""){
			sdk.urlHash(str).then(function(data){
				if(data["data"]["status"] == "success"){
					var url = "/news/s-"+data["data"]["url"]+"/";
					window.location.href = url;
				}
			});
		}
	}
});
</script>
`

func BenchmarkExtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extract(testSource)
	}
}
