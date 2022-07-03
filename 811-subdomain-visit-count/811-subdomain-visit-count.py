class Solution:
    def subdomainVisits(self, cpdomains: List[str]) -> List[str]:
        count=collections.Counter()
        for domain in cpdomains:
            n,s= domain.split()
            
            count[s]+=int(n)
            for i in range(len(s)):
                if s[i] == '.':
                    count[s[i + 1:]] += int(n)
        return ["%d %s" % (count[k], k) for k in count]
        