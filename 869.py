class Solution:
    def reorderedPowerOf2(self, n: int) -> bool:
        # Bizga kerakli n ni chegatasini aniqlash uchun
        limit = pow(10, 10)
        # Hamma 2 ni darajalari uchun list ochish
        ls = [1,]
        num = 2
        # listga solib borish
        while num < limit:
            ls.append(num)
            num *= 2
        
        # hamma kambinatsiyalarni tekshirish
        # n sonini raqamlarini hammasi listdagi 2 ni darajalarida
        # qatnashganligiga qarab tekshirish Misol uchun: 526 sonida
        # 256 (2 ni darajasi) bor hamma raqamlar bor
        n = list(str(n))
        n.sort()
        for i in ls:
            daraja = sorted(list(str(i)))
            if n == daraja:
                return True
        return False
    
print(Solution().reorderedPowerOf2(562))