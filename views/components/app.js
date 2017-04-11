let characters = [
    {name: "Daniel"},
    {name: "Gabriel"}
];

class Attribute extends React.Component{
    render(){
        return(
            <li>{this.props.name}</li>
        )
    }
}

class Character extends React.Component{
    render(){
        return(
            <ul>
                {this.props.characters.map ( character => {
                    return (
                        <Attribute name={character.name} />
                    )
                })}                
            </ul>
        )
    }
}

ReactDOM.render(<Character characters={characters} />, 
    document.getElementById('app'));