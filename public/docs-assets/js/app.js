$(function() {
	$('.bitcoin-address').popover({
		html: true
	});
	//回到顶部
	$.goup({
		trigger: 100,
		bottomOffset: 150,
		locationOffset: 100,
		title: '回到顶部',
		titleAsText: true
	});

})