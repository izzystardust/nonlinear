#! /usr/bin/python
import csv
from matplotlib import pyplot

prefixes = ["bisection", "chord", "newton", "secant", "shamanskii"]
suffixes = ['A', 'B', 'C', 'D', 'E', 'F']
colors = ('b', 'g', 'r', 'c', 'k')
linestyles = ['--', '-.', '-', ':', '-']
i = 0
for suffix in suffixes:
	residuals = []
	fig = pyplot.figure()
	fig.suptitle('Function ' + suffix)
	ax = fig.add_subplot(111)
	ax.set_xlabel('Iteration')
	ax.set_ylabel('Function value')
	for prefix in prefixes:
		with open('results/'+prefix+'-'+suffix+'.csv', 'rb') as csvfile:
			rdr = csv.reader(csvfile)
			column = []
			for row in rdr:
				column.append(row[1])
		linewidth = 2 if i == 4 else 1
		line, = ax.plot(map(abs, map(float, column[:100])), color='k', label=prefix, ls=linestyles[i], lw=linewidth)
		i = i+1
		if i > 4: i = 0
	ax.set_yscale('log')
	ax.legend()
	pyplot.show()
