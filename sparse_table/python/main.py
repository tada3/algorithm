# Implement the 'minmium' case here, the case of 'maximum' is almost the same.
class SparseTable:
	def __init__(self, a):
		self.a = a
		size = len(self.a)
		self.logs = self.build_logs(size)
		log_size = self.logs[size] + 1
		self.st = self.build_st(size, log_size)

	# st[i][j]: index corresponding to the minimum value in [j, j + 2**i) of a.
	# Here we use half open interval.
	# It is possible to set the minimum value to st[i][j], but we set an index because sometimes 
	# we need to get other attributes such as node ID.
	def build_st(self, n, m):
		st = [ [0]*n for i in range(m) ]

		# st[0][j] is an index of a that has the minimum value in [j, j+1).
		for j in range(n):
			st[0][j] = j
		
		for i in range(1, m):
			for j in range(n):
				end = j + (1 << i)
				# Value is not needed if end is out of range.
				if end > n:
					break
				st[i][j] = self.min_a(st[i-1][j], st[i-1][j+(1<<(i-1))])
		return st
		
	def query(self, x, y):
		if x == y:
			return x
		if x > y:
			x, y = y, x
		# Make it half open interval.
		y += 1
		log_len = self.logs[y-x]
		return self.min_a(self.st[log_len][x], self.st[log_len][y-(1<<log_len)])
	
	def min_a(self, x, y):
		if self.a[x] < self.a[y]:
			return x
		else:
			return y
	
	@staticmethod
	def build_logs(n):
		# logs[n] is used.
		logs = [0] * (n + 1)
		for i in range(2, n + 1):
			logs[i] = logs[i >> 1] + 1
		return logs


print('START')
a = [3, 2, 4, 9, 1, 5, 8]
print('a', a)
st = SparseTable(a)
min = st.query(1, 6)
print('min[1, 6]', min)
min = st.query(2, 5)
print('min[2, 5]', min)
min = st.query(3, 3)
print('min[3, 3]', min)
min = st.query(0, 6)
print('min[0, 6]', min)

a = [3, 2, 4, 9, 1, 5, 8, 0, 10, 500]
print('a', a)
st = SparseTable(a)
min = st.query(0, 8)
print('min[0, 8]', min)

a = [4]
print('a', a)
st = SparseTable(a)
min = st.query(0, 0)
print('min[0, 0]', min)

print('END')

