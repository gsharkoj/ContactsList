var SearchExample = React.createClass({

    getInitialState: function(){
        return { searchString: '' };
    },

	// получаем данные с сервера
	componentDidMount: function() {
		var _this = this;
		$.get('/data').then(function(response) {
		_this.setState({users: response})
		});
	},
  
    handleChange: function(e){
        this.setState({searchString:e.target.value});
    },
	
	// звонок	
	handleClick: function(i) {
	     
		alert("Вызов номера:" + i);
	},
  	
	// удаление контакта из списка
	handleClickDel: function(i) {	     	
	$.post("/del", { id: i} ).done(function( data ) {
    document.location.reload(true);
	});
  
	},
    render: function() {

        var libraries = this.props.items,
            searchString = this.state.searchString.trim().toLowerCase();

		if (this.state.users != undefined){
			libraries = this.state.users;
		}

        if(searchString.length > 0){

            libraries = libraries.filter(function(l){
                return l.Name.toLowerCase().match( searchString );
            });

        }

		// выводим весь список
		var inputStyle = {
					width: "100%"
				};
        return 	<div>
					<p className="text-center">						
					</p>
					<table className="table table-striped" id="Employees">
					<thead>
					<tr>
					<th>Сотрудники</th>
					<th>Телефон</th>
					</tr>
					</thead>					
					<tbody>		
					<tr>
					<td> <input style={inputStyle} type="text" value={this.state.searchString} onChange={this.handleChange} placeholder="Поиск по ФИО" /> </td>
					<td> </td>
					</tr>					
					{libraries.map(function(l){
							return <tr className = {l.Color}> 
									<td  width="80%">{l.Name}</td> 
									<td><button className="btn btn-primary" onClick={this.handleClick.bind(this,l.Url)}>{l.Url}</button>
									<button className="btn btn-danger" onClick={this.handleClickDel.bind(this,l.Id)}>-</button></td>
								   </tr>
                        }.bind(this)) }
					</tbody>		
					</table>
				</div>;		
		   
    }
});

                                                                                                                                                             
var elements = [];

React.render(
  <SearchExample items={elements}/>,
  document.getElementById('example')
);