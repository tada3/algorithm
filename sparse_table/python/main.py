# ここでは最小値の場合を実装する。最大値の場合もほぼ同じ。
class SparseTable:
	def __init__(self, a):
		self.a = a
		size = len(self.a)
		self.logs = self.build_logs(size)
		log_size = self.logs[size] + 1
		self.st = self.build_st(size, log_size)

	# st[i][j]: [j, j + 2**i) の最小値に対応するaのindex
	# 最小値そのものを値にしても良いが、他の値（ノードIDとか）が知りたい場合もあるので、indexにしておく
	def build_st(self, n, m):
		st = [ [0]*n for i in range(m) ]

		# st[0][j]は[j, j+1)の解に対応するindex、すなわちj
		for j in range(n):
			st[0][j] = j
		
		for i in range(1, m):
			for j in range(n):
				end = j + (1 << i)
				# endが範囲外の場合は必要ないので無視。
				if end > n:
					break
				st[i][j] = self.min_a(st[i-1][j], st[i-1][j+(1<<(i-1))])
		return st
		
	def query(self, x, y):
		if x == y:
			return x
		if x > y:
			x, y = y, x
		# 開区間にする [x, y)
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
		# longs[n] is used.
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

print('END')

