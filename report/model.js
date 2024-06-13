var colModel = [
	{
		title: "Name",
		width: 220,
		dataType: "string",
		dataIndx: "name",
	}, {
		title: "Total Iterations",
		width: 100,
		dataType: "integer",
		dataIndx: "total_iter_count",
	}, {
		title: "Marshal Iterations",
		width: 100,
		dataType: "integer",
		dataIndx: "marshal_iter_count",
		hidden: true,
	},  {
		title: "Unmarshal Iterations",
		width: 100,
		dataType: "integer",
		dataIndx: "unmarshal_iter_count",
		hidden: true,
	},  {
		title: "Total ns/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "total_ns_op",
	}, {
		title: "Marshal ns/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "marshal_ns_op",
		hidden: true,
	},  {
		title: "Unmarshal ns/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "unmarshal_ns_op",
		hidden: true,
	}, {
		title: "Total B/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "total_alloc_bytes",
	}, {
		title: "Marshal B/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "marshal_alloc_bytes",
		hidden: true,
	},  {
		title: "Unmarshal B/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "unmarshal_alloc_bytes",
		hidden: true,
	}, {
		title: "Total allocs/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "total_allocs",
	}, {
		title: "Marshal allocs/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "marshal_allocs",
		hidden: true,
	}, {
		title: "Unmarshal allocs/Op",
		width: 100,
		dataType: "integer",
		dataIndx: "unmarshal_allocs",
		hidden: true,
	}, {
		title: "Serialization Size",
		width: 100,
		dataType: "integer",
		dataIndx: "serialization_size",
	},  {
		title: "Unsafe Str Unmarshal",
		width: 100,
		type: "checkbox",
		dataIndx: "unsafe_string_unmarshal",
	}, {
		title: "Time Support",
		width: 100,
		type: "string",
		dataIndx: "time_support",
		filter: {
			type: "select",
		        condition: 'equal',
		        prepend: { '': '--Select--'  },
		        valueIndx: "time_support",
		        labelIndx: "time_support",
		        listeners: ['change']
		},
	},  {
		title: "API Kind",
		width: 100,
		type: "string",
		dataIndx: "api_kind",
		filter: {
			type: "select",
		        condition: 'equal',
		        prepend: { '': '--Select--'  },
		        valueIndx: "api_kind",
		        labelIndx: "api_kind",
		        listeners: ['change']
		},
	}, {
		title: "URL",
		width: 300,
		type: "string",
		dataIndx: "url",
		render: function (ui) {
			var url= ui.rowData.url;
			return "<a target=_blank href=\"https://"+url+"\"+>"+url+"</a>";
		},
	}, {
		title: "Notes",
		width: 300,
		type: "string",
		dataIndx: "notes",
	}
];
