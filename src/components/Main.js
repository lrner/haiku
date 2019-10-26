import React from 'react';
import axios from 'axios';

// import { ProgressBar } from 'react-bootstrap';

class Main extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      progress: 0,
      haiku: [],
      page: 1,
      building: 1,
      room: 1,
      row: 1,
      shelf: 1,
      series: 1,
      volume: 1,
      book: 1,
    }

    this.lastPage = this.lastPage.bind(this);
    this.nextPage = this.nextPage.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.changeBuilding = this.changeBuilding.bind(this)
    this.changeRoom = this.changeRoom.bind(this)
    this.changeRow = this.changeRow.bind(this)
    this.changeShelf = this.changeShelf.bind(this)
    this.changeSeries = this.changeSeries.bind(this)
    this.changeBook = this.changeBook.bind(this)
    this.changePage = this.changePage.bind(this)
  }

  nextPage() {
    this.setState({
      page: this.state.page + 1
    })

    this.loadPage();

  }

  lastPage() {
    this.setState({
      page: this.state.page - 1
    })

    this.loadPage();
  }

  loadPage() {
    var component = this;
    axios.post(`/book`, {
      building: 1,
      room: 1,
      row: 1,
      shelf: 1,
      series: 1,
      book: 1,
      page: 1,
    }).then(function(res) {
      component.setState({
        haiku: res.data
      });
    });
  }

  handleSubmit(ev) {
    ev.preventDefault();
    var component = this;
    axios.post(`/book`, {
      building: parseInt(this.state.building),
      room: parseInt(this.state.room),
      row: parseInt(this.state.row),
      shelf: parseInt(this.state.shelf),
      series: parseInt(this.state.series),
      book: parseInt(this.state.book),
      page: parseInt(this.state.page),
    }).then(function(res) {
      component.setState({
        haiku: res.data
      });
    });
  }

  componentDidMount() {
    this.loadPage()
  }

  changeBuilding(event) {
    this.setState({building: event.target.value});
  }

  changeRoom(event) {
    this.setState({room: event.target.value});
  }

  changeRow(event) {
    this.setState({row: event.target.value});
  }

  changeShelf(event) {
    this.setState({shelf: event.target.value});
  }

  changeSeries(event) {
    this.setState({series: event.target.value});
  }

  changeBook(event) {
    this.setState({book: event.target.value});
  }

  changePage(event) {
    this.setState({page: event.target.value});
  }

  render() {
    var haikuList = [];
    var x = 0;
    this.state.haiku.forEach((h) => {
      haikuList.push(
        <li key={x}>
          {h[0]}<br />
          {h[1]}<br />
          {h[2]}<br />
          <br />
        </li>
      );
      x++;
    });
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
        <div className="form-group">
          Building: <input type="number" value={this.state.building} onChange={this.changeBuilding} name="building" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Room: <input type="number"  value={this.state.room}        onChange={this.changeRoom} name="room" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Row: <input type="number"  value={this.state.row}          onChange={this.changeRow} name="room" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Shelf: <input type="number"  value={this.state.shelf} onChange={this.changeShelf} name="room" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Series: <input type="number"  value={this.state.series} onChange={this.changeSeries} name="room" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Book: <input type="number"  value={this.state.book} onChange={this.changeBook} name="room" min="1" max="1000000"/><br />
        </div>
        <div className="form-group">
          Page: <input type="number"  value={this.state.page} onChange={this.changePage} name="room" min="1" max="1000000"/><br />
        </div>
        <input type="submit" value="Submit" />
        </form>
        <ul>
          {haikuList}
        </ul>
      </div>
    );
  }
}

export default Main
