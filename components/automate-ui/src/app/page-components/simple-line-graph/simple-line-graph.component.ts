import {
  Component, Input, OnChanges, ViewChild, ElementRef, OnInit
} from '@angular/core';
import * as d3 from 'd3';
// import * as moment from 'moment';
// import { DateTime } from 'app/helpers/datetime/datetime';


@Component({
  selector: 'app-simple-line-graph',
  templateUrl: './simple-line-graph.component.html',
  styleUrls: ['./simple-line-graph.component.scss']
})

export class SimpleLineGraphComponent implements OnChanges, OnInit {

  constructor(
    private chart: ElementRef
  ) {}

  @Input() data: any = [];
  @Input() width = 900;
  @Input() height = 300;

  @ViewChild('svg', {static: true}) svg: ElementRef;

  ////////   X AXIS ITEMS   ////////
  // maps all of our x data points
  get xData() {
    return this.data.map(d => d.daysAgo);
  }
  // determines how wide the graph should be to hold our data
  // in its respective area;
  get rangeX() {
    const min = 50;
    const max = this.width - 50;
    return [min, max];
  }
  // determines the min and max values of the x axis
  get domainX() {
    const min = Math.min(...this.xData);
    const max = Math.max(...this.xData);
    return [min, max];
  }
  // determines each of our X axis points using the height and width of the chart
  get xScale() {
    return d3.scaleTime()
      .range(this.rangeX)
      .domain(this.domainX);
  }


  ////////   Y AXIS ITEMS   ////////

  // maps all of our Y data points
  get yData() {
    return this.data.map(d => d.percentage);
  }

  get rangeY() {
    const min = 10;
    const max = this.height - 10;
    return [max, min];
  }

  get domainY() {
    const min = 0;
    const max = 100; // since this based on a percentage do we want 0 to 100?
    return [min, max];
  }
  // determines each of our Y axis points using the height and width of the chart
  get yScale() {
    return d3.scaleLinear()
             .range(this.rangeY)
             .domain(this.domainY);
  }

  //////// SELECTIONS ////////
  get axisYSelection() {
    return this.svgSelection.select('.y-axis');
  }

  get svgSelection() {
    return d3.select(this.svg.nativeElement);
  }

  // returns a function that when passed our data, will return an svg path
  get createPath() {
    return d3.line()
              .x(d => this.xScale( d.daysAgo) )
              .y(d => this.yScale( d.percentage) );
  }


  get viewBox() {
    return `0 0 ${this.width} ${this.height}`;
  }

  resizeChart() {
    this.width = this.chart.nativeElement.getBoundingClientRect().width;
  }

    ////////////////// RENDER FUNCTIONS ////////////////////
  renderChart() {
    this.resizeChart();
    this.renderGrid();
    this.renderLine();
    this.renderPoints();
  }

  renderLine() {
    // create the line using path function
    const line = this.createPath(this.data);

    const theLine = this.svgSelection.selectAll('.line').data([this.data]);
    theLine.exit().remove();
    theLine.enter().append('path').attr('class', 'line').merge(theLine)
    .transition().duration(1000)
    .attr('d', line);
  }

  renderPoints() {
    const points = this.svgSelection.selectAll('circle').data(this.data);
    points.exit().remove();
    points
      .enter()
      .append('circle')
      .attr('class', 'circle')
      .merge(points)
      .transition().duration(1000)
      .attr('cx', d => this.xScale(d.daysAgo))
      .attr('cy', d => this.yScale(d.percentage))
      .attr('r', 4)
      .style('fill', 'black');
  }

  renderGrid() {
    // create the Y axis
    const yAxis = d3.axisRight(this.yScale)
                    .tickFormat(d => d + '%')
                    .ticks(1);
    // render the Y axis
    const y = this.svgSelection.selectAll('.y-axis').data([this.data]);
    y.exit().remove();
    y.enter().append('g').attr('class', 'y-axis').merge(y)
      .transition().duration(1000)
      .call(yAxis);

    // create the X axis grid lines
    const xGrid = d3.axisTop()
      .ticks(this.data.length)
      .tickFormat('')
      .tickSize(this.height - 20) // will need to subtract extra height here
      .tickSizeOuter(0)
      .scale(this.xScale);
    // Render the X axis and X ticks
    const grid = this.svgSelection.selectAll('.grid').data([this.data]);
    grid.exit().remove();
    grid.enter().append('g').attr('class', 'grid')
      // this line will need to be updated to flexible
      .attr('transform', `translate(0, ${this.height - 10})`)
      .merge(grid).transition().duration(1000)
      .call(xGrid);

    // remove zero from bottom of chart on x axis
    this.svgSelection.selectAll('.tick')
      .filter(tick => tick === 0)
      .remove();
  }


  ngOnChanges() {
    console.log('change');
    this.renderChart();
  }

  ngOnInit() {
    // console.log(this.data);
  }

  onResize() {
    this.resizeChart();
  }
}

