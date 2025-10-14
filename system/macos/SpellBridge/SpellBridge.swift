

import Foundation
import Cocoa

@_cdecl("CheckSpelling")
public func CheckSpelling(_ cstr: UnsafePointer<CChar>) -> UnsafePointer<CChar>? {
        let word = String(cString: cstr)
        let spellChecker = NSSpellChecker.shared
        let range = spellChecker.checkSpelling(of: word, startingAt: 0)

        if range.location == NSNotFound {
                return nil
            }

            
    }
